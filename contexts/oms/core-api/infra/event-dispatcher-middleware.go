package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type EventHandlerDispatcherMiddleware[T ddd.IBaseUseCase, R any] struct {
	orderLineService orderlines.IOrderLineService
}

func NewEventHandlerDispatcherMiddleware[T ddd.IBaseUseCase, R any](orderLineService orderlines.IOrderLineService) *EventHandlerDispatcherMiddleware[T, R] {
	return &EventHandlerDispatcherMiddleware[T, R]{
		orderLineService: orderLineService,
	}
}

func (s *EventHandlerDispatcherMiddleware[T, R]) Before(ctx context.Context, useCase T) (error, context.Context, T) {
	ctx = ddd.NewEventDispatcher(ctx, NewEventHandlerDispatcher(s.orderLineService))
	return nil, ctx, useCase
}
func (s *EventHandlerDispatcherMiddleware[T, R]) After(ctx context.Context, useCase T, err error, result *ddd.UseCaseResult[R]) (error, context.Context, T, *ddd.UseCaseResult[R]) {
	return nil, ctx, useCase, result
}
