package clients

import (
	"payment-service/clients/config"
	clients "payment-service/clients/invoice"
	config2 "payment-service/config"
)

type ClientRegistry struct{}

type IClientRegistry interface {
	GetInvoice() clients.IInvoiceClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientRegistry{}
}

func (c *ClientRegistry) GetInvoice() clients.IInvoiceClient {
	return clients.NewInvoiceClient(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.Invoice.Host),
		))
}
