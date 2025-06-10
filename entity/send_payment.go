package entity

import "encoding/xml"

// SendPaymentRequest — структура для парсинга входящего запроса
type SendPaymentRequest struct {
	XMLName            xml.Name `xml:"Request"`
	Id                 string   `xml:"Id,attr"`
	Service            string   `xml:"Service,attr"`
	Time               string   `xml:"Time,attr"`
	Value              string   `xml:"Value,attr"`
	Commission         string   `xml:"Commission,attr"`
	Currency           string   `xml:"Currency,attr"`
	PaymentTool        string   `xml:"PaymentTool,attr"`
	PaymentToolSubType string   `xml:"PaymentToolSubType,attr"`
	Session            string   `xml:"Session,attr"`

	PaymentParameters []struct {
		Name  string `xml:"Name,attr"`
		Value string `xml:"Value,attr"`
	} `xml:"PaymentParameters>Parameter"`
}

// SendPaymentResponse — структура для формирования ответа
type SendPaymentResponse struct {
	XMLName             xml.Name `xml:"Response"`
	PaymentId           string   `xml:"PaymentId,attr"`
	PaymentTime         string   `xml:"PaymentTime,attr"`
	State               string   `xml:"State,attr"`
	StateComment        string   `xml:"StateComment,attr"`
	ProcessingErrorCode string   `xml:"ProcessingErrorCode,attr"`
	Currency            string   `xml:"Currency,attr"`
}
