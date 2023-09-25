package ddd

import (
	"context"
)

type EventPublisherMiddleware[T IBaseUseCase, R any] struct {
	publisher IEventPublisher
}

func NewEventPublisherMiddleware[T IBaseUseCase, R any](publisher IEventPublisher) *EventPublisherMiddleware[T, R] {
	return &EventPublisherMiddleware[T, R]{
		publisher: publisher,
	}
}

func (s *EventPublisherMiddleware[T, R]) Before(ctx context.Context, useCase T) (error, context.Context, T) {
	return nil, ctx, useCase
}
func (s *EventPublisherMiddleware[T, R]) After(ctx context.Context, useCase T, err error, result *UseCaseResult[R]) (error, context.Context, T, *UseCaseResult[R]) {
	err = s.publisher.Publish(ctx)
	return err, ctx, useCase, result
}
