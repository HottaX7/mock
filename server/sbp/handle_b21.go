package sbp

import (
	"context"
	"encoding/xml"
	"espp-mock/entity"
	"espp-mock/metrics"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

func (h *Handler) HandleB21(w http.ResponseWriter, r *http.Request) {
	const op = "sbp.HandleB21"
	metrics.AddRequest("B21")

	// Получаем заголовки и данные из запроса
	correlationID := r.Header.Get("X-Correlation-ID")
	address := r.URL.Path

	// Извлекаем transactionNumber из URL (пример: /v1/C2BQRD/120000000020/B21)
	parts := strings.Split(address, "/")
	if len(parts) < 5 {
		log.Error().Str("op", op).Msg("Invalid URL format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	transactionNumber := parts[4]

	// Создаем контекст с метаданными
	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	ctx = context.WithValue(ctx, "transactionNumber", transactionNumber)

	// Читаем и парсим тело запроса
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Str("op", op).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var b21 entity.B21
	if err := xml.Unmarshal(bodyBytes, &b21); err != nil {
		log.Error().Err(err).Str("op", op).Str("body", string(bodyBytes)).Send()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверяем наличие Fr и To ID
	if b21.AppHdr.Fr.FIId.FinInstnId.Othr.ID == "" || b21.AppHdr.To.FIId.FinInstnId.Othr.ID == "" {
		log.Error().Str("op", op).Msg("Missing Fr or To ID in B21")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверяем наличие BizMsgIdr
	if b21.AppHdr.BizMsgIdr == "" {
		log.Error().Str("op", op).Msg("Missing BizMsgIdr in B21")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обрабатываем запрос через usecase
	if err := h.SBPUsecase.B21(ctx, &b21); err != nil {
		log.Error().Err(err).Str("op", op).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Формируем и отправляем ответ B22
	response := entity.NewB22Response(
		transactionNumber,
		b21.AppHdr.Fr.FIId.FinInstnId.Othr.ID, // fromId
		b21.AppHdr.To.FIId.FinInstnId.Othr.ID, // toId
		b21.AppHdr.BizMsgIdr,                  // bizMsgIdr
	)

	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("X-Sbp-Trn-Num", b21.AppHdr.BizMsgIdr)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Error().Err(err).Str("op", op).Msg("Failed to write response")
	}
}
