package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
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
	case "Orders:Created":
		return nil
	}

	// TODO: Add Outbox Pattern
	return nil
}
