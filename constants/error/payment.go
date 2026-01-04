package error

import "errors"

var (
	ErrInvoiceNotFound   = errors.New("invoice not found")
	ErrPaidAmountExceeds = errors.New("paid amount exceeds invoice amount")
)

var InvoiceErrors = []error{
	ErrInvoiceNotFound,
	ErrPaidAmountExceeds,
}
