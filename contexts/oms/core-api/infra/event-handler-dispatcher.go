package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/events/oms/orders"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type EventHandlerDispatcher struct {
	orderLineService orderlines.IOrderLineService
}

func NewEventHandlerDispatcher(orderLineService orderlines.IOrderLineService) ddd.IEventDispatcher {
	return &EventHandlerDispatcher{orderLineService: orderLineService}
}

func (s *EventHandlerDispatcher) Dispatch(ctx context.Context, event ddd.IBaseEvent) error {
	switch event.ExchangeName() {
	case "Orders:PaymentRejected":
		var h = orderlines.NewOrderPaymentRejectedEventHandler(s.orderLineService)

		return h.Handle(ctx, event.(*orders.PaymentRejected))
	}

	// TODO: Add Outbox Pattern
	return nil
}
