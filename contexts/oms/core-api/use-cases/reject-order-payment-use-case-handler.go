package use_cases

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orders"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/infra"
	orderModels "github.com/sercanakmaz/go-boilerplate-v3/models/order"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
)

type RejectOrderPaymentUseCaseHandler struct {
	orderService     orders.IOrderService
	orderLineService orderlines.IOrderLineService
	middlewares      []ddd.IBaseUseCaseMiddleware[*orderModels.RejectOrderPaymentCommand, string]
}

func NewRejectOrderPaymentUseCaseHandler(orderService orders.IOrderService, orderLineService orderlines.IOrderLineService) *RejectOrderPaymentUseCaseHandler {
	var handler = &RejectOrderPaymentUseCaseHandler{
		orderService:     orderService,
		orderLineService: orderLineService,
	}

	handler.middlewares = append(handler.middlewares, &mongo.TransactionMiddleware[*orderModels.RejectOrderPaymentCommand, string]{})
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
