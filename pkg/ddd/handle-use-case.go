package ddd

import (
	"context"
)

func HandleUseCase[H IBaseUseCaseHandler[U, R], U IBaseUseCase, R any](ctx context.Context, handler H, useCase U, result *UseCaseResult[R]) error {
	var (
		handleErr     error
		middlewareErr error
		dispatcherErr error
		middleWares   = handler.GetMiddlewares()
		innerResult   *UseCaseResult[R]
	)

	ctx = NewEventContext(ctx)

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase = middleWare.Before(ctx, useCase); middlewareErr != nil {
			return middlewareErr
		}
	}

	handleErr, innerResult = handler.Handle(ctx, useCase)

	*result = *innerResult

	dispatcher := GetEventDispatcher(ctx)

	if dispatcher != nil {
		eventContext := GetEventContext(ctx)

		for true {
			raisedEvent := eventContext.TakeRaised()
			if raisedEvent == nil {
				break
			}
			if dispatcherErr = dispatcher.Dispatch(ctx, raisedEvent); dispatcherErr != nil {
				return dispatcherErr
			}
			eventContext.AddDispatched(raisedEvent)
		}
	}

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase, result = middleWare.After(ctx, useCase, handleErr, result); middlewareErr != nil {
			return middlewareErr
		}
	}

	return handleErr
}
