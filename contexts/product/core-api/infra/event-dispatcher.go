package infra

import (
	"context"
	"fmt"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type EventDispatcher struct {
}

func NewEventDispatcher() ddd.IEventDispatcher {
	return &EventDispatcher{}
}

func (s *EventDispatcher) Dispatch(ctx context.Context, event ddd.IBaseEvent) error {
	switch event.EventName() {
	case "Product:Created":
		fmt.Println("Sync event dispatcher!")
		return nil
	}
	return nil
}
