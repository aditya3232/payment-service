package dto

type KafkaEvent struct {
	Name string `json:"name"`
}

type KafkaMetaData struct {
	Sender    string `json:"sender"`
	SendingAt string `json:"sending_at"`
}

type KafkaData struct {
	PaymentID   int     `json:"payment_id"`
	InvoiceID   int     `json:"invoice_id"`
	Amount      float64 `json:"amount"`
	ReferenceNo string  `json:"reference_no"`
}

type KafkaBody struct {
	Type string     `json:"type"`
	Data *KafkaData `json:"data"`
}

type KafkaMessage struct {
	Event    KafkaEvent    `json:"event"`
	Metadata KafkaMetaData `json:"metadata"`
	Body     KafkaBody     `json:"body"`
}
