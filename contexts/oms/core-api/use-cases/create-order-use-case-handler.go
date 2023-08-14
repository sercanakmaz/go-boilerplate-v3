package use_cases

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orders"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/infra"
	orderModels "github.com/sercanakmaz/go-boilerplate-v3/models/order"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourMongo "github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateOrderUseCaseHandler struct {
	client           *mongo.Client
	orderService     orders.IOrderService
	orderLineService orderlines.IOrderLineService
	middlewares      []ddd.IBaseUseCaseMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse]
}

func NewCreateOrderUseCaseHandler(client *mongo.Client, orderService orders.IOrderService, orderLineService orderlines.IOrderLineService) *CreateOrderUseCaseHandler {
	var handler = &CreateOrderUseCaseHandler{
		client:           client,
		orderService:     orderService,
		orderLineService: orderLineService,
	}

	handler.middlewares = append(handler.middlewares, ourMongo.NewTransactionMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse](client))
	handler.middlewares = append(handler.middlewares, infra.NewEventHandlerDispatcherMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse](orderLineService))

	return handler
}

func (self *CreateOrderUseCaseHandler) GetMiddlewares() []ddd.IBaseUseCaseMiddleware[*orderModels.CreateOrderCommand, *orderModels.CreateOrderResponse] {
	return self.middlewares
}

func (self *CreateOrderUseCaseHandler) Handle(ctx context.Context, command *orderModels.CreateOrderCommand) (error, *ddd.UseCaseResult[*orderModels.CreateOrderResponse]) {

	var (
		err        error
		order      *orders.Order
		orderLines []*orderlines.OrderLine
	)

	if order, err = self.orderService.AddNew(ctx,
		command.OrderNumber,
		command.Price); err != nil {
		return err, nil
	}

	for _, line := range command.OrderLines {
		var orderLine *orderlines.OrderLine

		if orderLine, err = self.orderLineService.AddNew(ctx,
			line.Sku,
			command.OrderNumber,
			line.Price); err != nil {
			return err, nil
		}

		orderLines = append(orderLines, orderLine)
	}

	return nil, ddd.NewUseCaseResultWithContent[*orderModels.CreateOrderResponse](orderModels.NewCreateOrderResponse(order, orderLines))
}
