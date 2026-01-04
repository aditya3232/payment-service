package repositories

import (
	paymentRepo "payment-service/repositories/payment"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetPayment() paymentRepo.IPaymentRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetPayment() paymentRepo.IPaymentRepository {
	return paymentRepo.NewPaymentRepository(r.db)
}
