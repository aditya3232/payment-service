package services

import (
	"context"
	"payment-service/clients"
	errConstant "payment-service/constants/error"
	"payment-service/domain/dto"
	"payment-service/repositories"
)

type PaymentService struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
}

type IPaymentService interface {
	Create(context.Context, *dto.PaymentRequest) (*dto.PaymentResponse, error)
	FindAllWithoutPagination(context.Context, *dto.PaymentRequestParam) ([]dto.PaymentResponse, error)
}

func NewPaymentService(repository repositories.IRepositoryRegistry, client clients.IClientRegistry) IPaymentService {
	return &PaymentService{repository: repository, client: client}
}

func (s *PaymentService) Create(ctx context.Context, req *dto.PaymentRequest) (*dto.PaymentResponse, error) {
	_, err := s.client.GetInvoice().FindByID(ctx, req.InvoiceID)
	if err != nil {
		return nil, errConstant.ErrInvoiceNotFound
	}

	payment, err := s.repository.GetPayment().Create(ctx, &dto.PaymentRequest{
		InvoiceID:   req.InvoiceID,
		Amount:      req.Amount,
		Method:      req.Method,
		ReferenceNo: req.ReferenceNo,
		PaidAt:      req.PaidAt,
	})

	if err != nil {
		return nil, err
	}

	response := &dto.PaymentResponse{
		ID:          payment.ID,
		InvoiceID:   payment.InvoiceID,
		Amount:      payment.Amount,
		Method:      payment.Method,
		ReferenceNo: payment.ReferenceNo,
		PaidAt:      payment.PaidAt,
		CreatedAt:   payment.CreatedAt,
	}

	return response, nil
}

func (s *PaymentService) FindAllWithoutPagination(ctx context.Context, req *dto.PaymentRequestParam) ([]dto.PaymentResponse, error) {
	payments, err := s.repository.GetPayment().FindAllWithoutPagination(ctx, req)
	if err != nil {
		return nil, err
	}

	paymentResults := make([]dto.PaymentResponse, 0, len(payments))
	for _, payment := range payments {
		paymentResults = append(paymentResults, dto.PaymentResponse{
			ID:          payment.ID,
			InvoiceID:   payment.InvoiceID,
			Amount:      payment.Amount,
			Method:      payment.Method,
			ReferenceNo: payment.ReferenceNo,
			PaidAt:      payment.PaidAt,
			CreatedAt:   payment.CreatedAt,
		})
	}

	return paymentResults, nil
}
