package clients

import (
	"context"
	"fmt"
	"net/http"
	"payment-service/clients/config"
)

type InvoiceClient struct {
	client config.IClientConfig
}

type IInvoiceClient interface {
	FindByID(context.Context, int) (*InvoiceData, error)
}

func NewInvoiceClient(client config.IClientConfig) IInvoiceClient {
	return &InvoiceClient{client: client}
}

func (u *InvoiceClient) FindByID(ctx context.Context, id int) (*InvoiceData, error) {
	var response InvoiceResponse
	request := u.client.Client().
		Get(fmt.Sprintf("%s/api/v1/invoices/%d", u.client.BaseURL(), id))

	resp, _, errs := request.EndStruct(&response)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invoice response: %s", response.Message)
	}

	return &response.Data, nil
}
