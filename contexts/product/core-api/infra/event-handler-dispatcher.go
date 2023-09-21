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

// TODO: Sercan'a sor.
func (s *EventHandlerDispatcher) Dispatch(ctx context.Context, event ddd.IBaseEvent) error {
	switch event.EventName() {
	case "Product:Created":
		fmt.Println("Sync event dispatcher!")
		return nil
	}
	return nil
}
