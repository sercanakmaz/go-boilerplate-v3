package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
)

type EventDispatcher struct {
	Logger logger.Logger
}

func NewEventDispatcher(log logger.Logger) ddd.IEventDispatcher {
	return &EventDispatcher{Logger: log}
}

func (s *EventDispatcher) Dispatch(ctx context.Context, event ddd.IBaseEvent) error {
	switch event.EventName() {
	case "Product:Created":
		s.Logger.Info(ctx, "Sync event processed!")
		return nil
	}
	return nil
}
