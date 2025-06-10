package sbp

import (
	"encoding/json"
	"espp-mock/metrics"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (h *Handler) HandlePaymentData(w http.ResponseWriter, r *http.Request) {
	uniQrID := chi.URLParam(r, "id")         // Получаем {id} из маршрута
	amountStr := r.URL.Query().Get("amount") // Получаем amount из query-параметров

	metrics.AddRequest("C2BQRD")

	// Логируем значение amountStr перед обработкой
	log.Info().Str("amountStr", amountStr).Msg("Received amount before processing")

	// Если amount не передан, задаем значение по умолчанию
	if amountStr == "" {
		amountStr = "1100"
	}

	// Заменяем запятую на точку, если это необходимо
	amountStr = strings.ReplaceAll(amountStr, ",", ".")

	// Логируем значение amountStr после замены запятой на точку
	log.Info().Str("amountStr", amountStr).Msg("Received amount after processing")

	// Проверяем, что amount является корректным числом
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Error().Err(err).Str("amountStr", amountStr).Msg("Failed to parse amount")
		http.Error(w, `{"error": "invalid amount format"}`, http.StatusBadRequest)
		return
	}

	// Логируем успешное преобразование
	log.Info().Str("uniQrId", uniQrID).Float64("amount", amount).Msg("Successfully parsed amount")

	// Формируем JSON-ответ
	response := map[string]interface{}{
		"code":    "RQ00000",
		"message": "Запрос обработан успешно",
		"data": map[string]interface{}{
			"uniQrId": uniQrID,
			/*"uniQrId":        "BS10006LQEGL0BL78D9R7SO6O93SGJVN", //жестко зашиваем QR*/
			"uniQrType":      "02",
			"scenario":       "C2B",
			"legalName":      "ООО \"ТД АМТ\"",
			"memberId":       "120000000020",
			"brandName":      "тест_2",
			"amount":         amount, // Отправляем amount (либо переданный, либо дефолтный)
			"paymentPurpose": "Rath, Oberbrunner and Howell",
			"address":        "Уфа",
			"mcc":            "5047",
			"ogrn":           "1200200041360",
			"inn":            "0277950370",
			"redirectUrl":    nil,
			"agentId":        "A11000000072",
			"merchantId":     "MA0000018139",
			"countryCode":    "RU",
			"receiverBic":    "044525111",
			"responseId":     "79f9c5fe3af74300b6dd13fea2c09d27",
			"paramsId":       nil,
			"paymentMethodData": []map[string]interface{}{
				{
					"paymentServiceId": "PS0000000001",
					"version":          "v1",
					"additionalData": map[string]string{
						"fraudScore":      "FFFFFDDDDDDDDDDD",
						"receiverAccount": "40702810041000001132",
					},
				},
			},
		},
	}

	// Преобразуем в JSON и отправляем
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
