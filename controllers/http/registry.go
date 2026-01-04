package controllers

import (
	paymentController "payment-service/controllers/http/payment"
	"payment-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetPayment() paymentController.IPaymentController
}

func NewControllerregistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetPayment() paymentController.IPaymentController {
	return paymentController.NewPaymentController(r.service)
}
