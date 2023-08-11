package orderlines

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type (
	IOrderLineService interface {
		AddNew(ctx context.Context, sku, orderNumber string, price shared.Money) (*OrderLine, error)
		GetBySku(ctx context.Context, sku string) (*OrderLine, error)
	}
	orderLineService struct {
		Repository IOrderLineRepository
	}
)

func NewOrderLineService(repository IOrderLineRepository) IOrderLineService {
	return &orderLineService{Repository: repository}
}

func (service orderLineService) AddNew(ctx context.Context, sku, orderNUmber string, price shared.Money) (*OrderLine, error) {
	var orderLine = NewOrderLine(sku, orderNUmber, price)

	var err = service.Repository.Add(ctx, orderLine)

	return orderLine, err
}

func (service orderLineService) GetBySku(ctx context.Context, sku string) (*OrderLine, error) {
	var (
		err       error
		orderLine *OrderLine
	)

	if orderLine, err = service.Repository.FindOneBySku(ctx, sku); err != nil {
		return nil, err
	}

	return orderLine, err
}
