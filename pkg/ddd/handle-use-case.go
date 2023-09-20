package ddd

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
)

func HandleUseCase[H IBaseUseCaseHandler[U, R], U IBaseUseCase, R any](ctx context.Context, messageBus *rabbitmqv1.Client, handler H, useCase U, result *UseCaseResult[R]) error {
	var (
		handleErr     error
		middlewareErr error
		dispatcherErr error
		middleWares   = handler.GetMiddlewares()
		innerResult   *UseCaseResult[R]
	)

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase = middleWare.Before(ctx, useCase); middlewareErr != nil {
			return middlewareErr
		}
	}

	handleErr, innerResult = handler.Handle(ctx, useCase)

	if innerResult != nil {
		*result = *innerResult
	}

	if handleErr == nil {
		dispatcher := GetEventDispatcher(ctx)

		if dispatcher != nil {
			eventContext := GetEventContext(ctx)

			for true {
				raisedEvent := eventContext.TakeRaised()
				if raisedEvent == nil {
					break
				}
				if dispatcherErr = dispatcher.Dispatch(ctx, raisedEvent); dispatcherErr != nil {
					handleErr = dispatcherErr
					break
				}
				eventContext.AddDispatched(raisedEvent)
			}
		}
	}

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase, result = middleWare.After(ctx, useCase, handleErr, result); middlewareErr != nil {
			return middlewareErr
		}
	}

	return handleErr
}
