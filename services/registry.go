package services

import (
	clients "payment-service/clients/midtrans"
	utilminio "payment-service/common/minio"
	"payment-service/controllers/kafka"
	"payment-service/repositories"
	services "payment-service/services/payment"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
	minio      utilminio.IMinioClient
	kafka      kafka.IKafkaRegistry
	midtrans   clients.IMidtransClient
}

type IServiceRegistry interface {
	GetPayment() services.IPaymentService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry, minio utilminio.IMinioClient, kafka kafka.IKafkaRegistry, midtrans clients.IMidtransClient) IServiceRegistry {
	return &Registry{
		repository: repository,
		minio:      minio,
		kafka:      kafka,
		midtrans:   midtrans,
	}
}

func (r *Registry) GetPayment() services.IPaymentService {
	return services.NewPaymentService(r.repository, r.minio, r.kafka, r.midtrans)
}
