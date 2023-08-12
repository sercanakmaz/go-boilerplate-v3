package orderlines

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IOrderLineService interface {
		AddNew(ctx context.Context, sku, orderNumber string, price shared.Money) (*OrderLine, error)
		GetById(ctx context.Context, sku string) (*OrderLine, error)
		FindByOrderNumber(ctx context.Context, orderNumber string) ([]*OrderLine, error)
		Cancel(ctx context.Context, id string, cancelReason string) (*OrderLine, error)
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

func (service orderLineService) Cancel(ctx context.Context, id string, cancelReason string) (*OrderLine, error) {
	var (
		err       error
		idTyped   primitive.ObjectID
		orderLine *OrderLine
	)

	if idTyped, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	if orderLine, err = service.Repository.FindOneById(ctx, idTyped); err != nil {
		return orderLine, err
	}

	orderLine.Cancel(cancelReason)

	err = service.Repository.Update(ctx, orderLine)

	return orderLine, err
}
func (service orderLineService) GetById(ctx context.Context, id string) (*OrderLine, error) {
	var (
		err       error
		idTyped   primitive.ObjectID
		orderLine *OrderLine
	)

	if idTyped, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	if orderLine, err = service.Repository.FindOneById(ctx, idTyped); err != nil {
		return nil, err
	}

	return orderLine, err
}

func (service orderLineService) FindByOrderNumber(ctx context.Context, orderNumber string) ([]*OrderLine, error) {
	var (
		err        error
		orderLines []*OrderLine
	)

	if orderLines, err = service.Repository.FindByOrderNumber(ctx, orderNumber); err != nil {
		return nil, err
	}

	return orderLines, err
}
