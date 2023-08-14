package ddd

import (
	"context"
)

type IEventDispatcher interface {
	Dispatch(ctx context.Context, event IBaseEvent) error
}

var eventDispatcherKey = "eventDispatcher"

func newEventDispatcher(ctx context.Context, dispatcher IEventDispatcher) context.Context {
	return context.WithValue(ctx, eventDispatcherKey, dispatcher)
}

func GetEventDispatcher(ctx context.Context) IEventDispatcher {
	var result = ctx.Value(eventDispatcherKey)

	if result == nil {
		return nil
	}

	return result.(IEventDispatcher)
}
