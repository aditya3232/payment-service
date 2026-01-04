package services

import (
	"context"
	"encoding/json"
	"fmt"
	"payment-service/clients"
	"payment-service/common/util"
	"payment-service/config"
	"payment-service/constants"
	errConstant "payment-service/constants/error"
	"payment-service/controllers/kafka"
	"payment-service/domain/dto"
	"payment-service/repositories"
	"time"
)

type PaymentService struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
	kafka      kafka.IKafkaRegistry
}

type IPaymentService interface {
	Create(context.Context, *dto.PaymentRequest) (*dto.PaymentResponse, error)
	FindAllWithoutPagination(context.Context, *dto.PaymentRequestParam) ([]dto.PaymentResponse, error)
}

func NewPaymentService(repository repositories.IRepositoryRegistry, client clients.IClientRegistry, kafka kafka.IKafkaRegistry) IPaymentService {
	return &PaymentService{repository: repository, client: client, kafka: kafka}
}

func (s *PaymentService) produceToKafka(req *dto.PaymentToEventRequest) error {
	event := dto.KafkaEvent{
		Name: constants.EventName,
	}

	metadata := dto.KafkaMetaData{
		Sender:    "payment-service",
		SendingAt: time.Now().Format(time.RFC3339),
	}

	body := dto.KafkaBody{
		Type: "JSON",
		Data: &dto.KafkaData{
			PaymentID:   req.PaymentID,
			InvoiceID:   req.InvoiceID,
			Amount:      req.Amount,
			ReferenceNo: req.ReferenceNo,
		},
	}

	kafkaMessage := dto.KafkaMessage{
		Event:    event,
		Metadata: metadata,
		Body:     body,
	}

	topic := config.Config.Kafka.Topic
	kafkaMessageJSON, _ := json.Marshal(kafkaMessage)
	key := []byte(fmt.Sprintf("invoice-%d", req.InvoiceID))

	err := s.kafka.GetKafkaProducer().ProduceMessage(topic, key, kafkaMessageJSON)
	if err != nil {
		return err
	}

	return nil
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
		PaidAt:      time.Now(),
	})

	// handle ketika duplicate key postgres reference_no (idempotent safe)
	if err != nil {
		if util.IsUniqueViolation(err) {
			p, err := s.repository.GetPayment().
				FindByReferenceNo(ctx, req.ReferenceNo)
			if err != nil {
				return nil, err
			}

			return &dto.PaymentResponse{
				ID:          p.ID,
				InvoiceID:   p.InvoiceID,
				Amount:      p.Amount,
				Method:      p.Method,
				ReferenceNo: p.ReferenceNo,
				PaidAt:      p.PaidAt,
				CreatedAt:   p.CreatedAt,
			}, nil
		}

		return nil, err
	}

	err = s.produceToKafka(&dto.PaymentToEventRequest{
		PaymentID:   payment.ID,
		InvoiceID:   payment.InvoiceID,
		Amount:      payment.Amount,
		ReferenceNo: payment.ReferenceNo,
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
