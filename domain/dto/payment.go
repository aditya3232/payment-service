package dto

import "time"

type PaymentRequest struct {
	InvoiceID   int     `json:"invoice_id" validate:"required,gt=0"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Method      string  `json:"method" validate:"required,gt=0"`
	ReferenceNo string  `json:"reference_no" validate:"required,gt=0"`
	PaidAt      string  `json:"paid_at" validate:"datetime=2006-01-02"`
}

type PaymentRequestParam struct {
	InvoiceID *int `form:"invoice_id"`
}

type PaymentResponse struct {
	ID          int64      `json:"id"`
	InvoiceID   int64      `json:"invoice_id"`
	Amount      float64    `json:"amount"`
	Method      string     `json:"method"`
	ReferenceNo string     `json:"reference_no"`
	PaidAt      time.Time  `json:"paid_at"`
	CreatedAt   *time.Time `json:"created_at"`
}
