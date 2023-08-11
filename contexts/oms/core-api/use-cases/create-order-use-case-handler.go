package use_cases

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orders"
	orderModels "github.com/sercanakmaz/go-boilerplate-v3/models/order"
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/use-case"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"net/http"
)

type CreateOrderUseCaseHandler struct {
	orderService     orders.IOrderService
	orderLineService orderlines.IOrderLineService
	middlewares      []use_case.IBaseUseCaseMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse]
}

func NewCreateOrderUseCaseHandler(orderService orders.IOrderService) *CreateOrderUseCaseHandler {
	var handler = &CreateOrderUseCaseHandler{
		orderService: orderService,
	}

	handler.middlewares = append(handler.middlewares, &mongo.TransactionMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse]{})

	return handler
}

func (self *CreateOrderUseCaseHandler) GetMiddlewares() []use_case.IBaseUseCaseMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse] {
	return self.middlewares
}

func (self *CreateOrderUseCaseHandler) Handle(ctx context.Context, command *orderModels.CreateOrderCommand) (error, *use_case.UseCaseResult[*orderModels.CreateOrderResponse]) {

	var (
		err        error
		order      *orders.Order
		orderLines []*orderlines.OrderLine
	)

	if order, err = self.orderService.AddNew(ctx,
		command.OrderNumber,
		command.Price); err != nil {
		return err, use_case.NewUseCaseResult[*orderModels.CreateOrderResponse](http.StatusInternalServerError)
	}

	for _, line := range command.OrderLines {
		var orderLine *orderlines.OrderLine

		if orderLine, err = self.orderLineService.AddNew(ctx,
			line.Sku,
			command.OrderNumber,
			line.Price); err != nil {
			return err, use_case.NewUseCaseResult[*orderModels.CreateOrderResponse](http.StatusInternalServerError)
		}

		orderLines = append(orderLines, orderLine)
	}

	return nil, use_case.NewUseCaseResultWithContent[*orderModels.CreateOrderResponse](http.StatusCreated, orderModels.NewCreateOrderResponse(order, orderLines))
}
