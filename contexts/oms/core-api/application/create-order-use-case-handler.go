package application

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/domain/orders"
	order2 "github.com/sercanakmaz/go-boilerplate-v3/models/order"
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/use-case"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"net/http"
)

type CreateOrderUseCaseHandler struct {
	orderService orders.IOrderService
	middlewares  []use_case.IBaseUseCaseMiddleware[*order2.CreateOrderCommand, *orders.Order]
}

func NewCreateOrderUseCaseHandler(orderService orders.IOrderService) *CreateOrderUseCaseHandler {
	var handler = &CreateOrderUseCaseHandler{
		orderService: orderService,
	}

	handler.middlewares = append(handler.middlewares, &mongo.TransactionMiddleware[*order2.CreateOrderCommand, *orders.Order]{})

	return handler
}

func (self *CreateOrderUseCaseHandler) GetMiddlewares() []use_case.IBaseUseCaseMiddleware[*order2.CreateOrderCommand, *orders.Order] {
	return self.middlewares
}

func (self *CreateOrderUseCaseHandler) Handle(ctx context.Context, command *order2.CreateOrderCommand) (error, *use_case.UseCaseResult[*orders.Order]) {

	var (
		err   error
		order *orders.Order
	)

	if order, err = self.orderService.AddNew(ctx,
		command.OrderNumber,
		command.Price); err != nil {
		return err, use_case.NewUseCaseResult[*orders.Order](http.StatusInternalServerError)
	}

	return nil, use_case.NewUseCaseResultWithContent[*orders.Order](http.StatusCreated, order)
}
