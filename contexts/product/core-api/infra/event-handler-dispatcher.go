package infra

import (
	"context"
	"fmt"
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
		fmt.Println("here")
		return nil
	}
	return nil
}
