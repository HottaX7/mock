// handle_b22.go

package sbp

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"espp-mock/entity"

	"github.com/rs/zerolog/log"
)

func (h *Handler) HandleB22(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleB22"
	logger := log.With().Str("op", op).Logger()

	// 1. Читаем тело запроса
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read B22 request body")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 2. Парсим в структуру entity.B22
	var b22 entity.B22
	if err := xml.Unmarshal(bodyBytes, &b22); err != nil {
		logger.Error().Err(err).Str("body", string(bodyBytes)).Msg("Failed to parse B22 XML")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 3. Получаем transactionNumber из контекста (если нужно для логики)
	transactionNumber := contextFromRequest(r).Value("transactionNumber").(string)

	// 4. Логируем успешный приём
	logger.Info().
		Str("transactionNumber", transactionNumber).
		Interface("request", b22).
		Msg("Received B22")

	// 5. Отвечаем шлюзу
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// contextFromRequest — вспомогательная функция получения контекста с ID транзакции
func contextFromRequest(r *http.Request) context.Context {
	correlationID := r.Header.Get("X-Correlation-ID")
	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	return ctx
}
