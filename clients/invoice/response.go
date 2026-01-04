package clients

import "time"

type InvoiceResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    InvoiceData `json:"data"`
}

type InvoiceData struct {
	ID         int        `json:"id"`
	CustomerID int        `json:"customer_id"`
	Amount     float64    `json:"amount"`
	PaidAmount float64    `json:"paid_amount"`
	Currency   string     `json:"currency"`
	DueDate    time.Time  `json:"due_date"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
