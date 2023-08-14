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

type RejectOrderPaymentUseCaseHandler struct {
	client           *mongo.Client
	orderService     orders.IOrderService
	orderLineService orderlines.IOrderLineService
	middlewares      []ddd.IBaseUseCaseMiddleware[*orderModels.RejectOrderPaymentCommand, string]
}

func NewRejectOrderPaymentUseCaseHandler(client *mongo.Client, orderService orders.IOrderService, orderLineService orderlines.IOrderLineService) *RejectOrderPaymentUseCaseHandler {
	var handler = &RejectOrderPaymentUseCaseHandler{
		client:           client,
		orderService:     orderService,
		orderLineService: orderLineService,
	}

	handler.middlewares = append(handler.middlewares, ourMongo.NewTransactionMiddleware[*orderModels.RejectOrderPaymentCommand, string](client))
	handler.middlewares = append(handler.middlewares, infra.NewEventHandlerDispatcherMiddleware[*orderModels.RejectOrderPaymentCommand, string](orderLineService))

	return handler
}

func (self *RejectOrderPaymentUseCaseHandler) GetMiddlewares() []ddd.IBaseUseCaseMiddleware[*orderModels.RejectOrderPaymentCommand, string] {
	return self.middlewares
}

func (self *RejectOrderPaymentUseCaseHandler) Handle(ctx context.Context, command *orderModels.RejectOrderPaymentCommand) (error, *ddd.UseCaseResult[string]) {

	var (
		err error
	)

	if _, err = self.orderService.RejectPayment(ctx,
		command.OrderNumber, command.RejectReason); err != nil {
		return err, nil
	}

	return nil, ddd.NewUseCaseResultWithContent[string]("")
}
