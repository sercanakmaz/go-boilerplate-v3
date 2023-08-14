package ddd

import (
	"context"
)

type EventHandlerDispatcherMiddleware[T IBaseUseCase, R any] struct {
	dispatcher IEventDispatcher
}

func NewEventHandlerDispatcherMiddleware[T IBaseUseCase, R any](dispatcher IEventDispatcher) *EventHandlerDispatcherMiddleware[T, R] {
	return &EventHandlerDispatcherMiddleware[T, R]{
		dispatcher: dispatcher,
	}
}

func (s *EventHandlerDispatcherMiddleware[T, R]) Before(ctx context.Context, useCase T) (error, context.Context, T) {
	ctx = newEventDispatcher(ctx, s.dispatcher)
	return nil, ctx, useCase
}
func (s *EventHandlerDispatcherMiddleware[T, R]) After(ctx context.Context, useCase T, err error, result *UseCaseResult[R]) (error, context.Context, T, *UseCaseResult[R]) {
	return nil, ctx, useCase, result
}
