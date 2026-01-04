package services

import (
	"payment-service/clients"
	"payment-service/controllers/kafka"
	"payment-service/repositories"
	paymentService "payment-service/services/payment"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
	kafka      kafka.IKafkaRegistry
}

type IServiceRegistry interface {
	GetPayment() paymentService.IPaymentService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry, client clients.IClientRegistry, kafka kafka.IKafkaRegistry) IServiceRegistry {
	return &Registry{repository: repository, client: client, kafka: kafka}
}

func (r *Registry) GetPayment() paymentService.IPaymentService {
	return paymentService.NewPaymentService(r.repository, r.client, r.kafka)
}
