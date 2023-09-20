package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type EventHandlerDispatcher struct {
}

func NewEventHandlerDispatcher() ddd.IEventDispatcher {
	return &EventHandlerDispatcher{}
}

func (s *EventHandlerDispatcher) Dispatch(ctx context.Context, event ddd.IBaseEvent) error {
	switch event.ExchangeName() {
	case "Product:Created":
		// TODO: Sercan'a sor -> Rabbit'e pushla ?
		return nil
	}

	// TODO: Add Outbox Pattern
	return nil
}
