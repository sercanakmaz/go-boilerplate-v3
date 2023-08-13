package orders

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type (
	IOrderService interface {
		AddNew(ctx context.Context, orderNumber string, price shared.Money) (*Order, error)
		RejectPayment(ctx context.Context, orderNumber string, paymentRejectReason string) (*Order, error)
		GetByOrderNumber(ctx context.Context, orderNumber string) (*Order, error)
	}
	orderService struct {
		Repository IOrderRepository
	}
)

func NewOrderService(repository IOrderRepository) IOrderService {
	return &orderService{Repository: repository}
}

func (service orderService) AddNew(ctx context.Context, orderNumber string, price shared.Money) (*Order, error) {
	var order = NewOrder(orderNumber, price)

	var err = service.Repository.Add(ctx, order)

	return order, err
}

func (service orderService) RejectPayment(ctx context.Context, orderNumber string, paymentRejectReason string) (*Order, error) {
	var (
		err   error
		order *Order
	)

	if order, err = service.Repository.FindOneByOrderNumber(ctx, orderNumber); err != nil {
		return order, err
	}

	order.RejectPayment(paymentRejectReason)

	err = service.Repository.Update(ctx, order)

	return order, err
}

func (service orderService) GetByOrderNumber(ctx context.Context, orderNumber string) (*Order, error) {
	var (
		err   error
		order *Order
	)

	if order, err = service.Repository.FindOneByOrderNumber(ctx, orderNumber); err != nil {
		return nil, err
	}

	return order, err
}
