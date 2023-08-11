package use_case

import (
	"context"
	event_handler "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/event-handler"
)

func Handle[H IBaseUseCaseHandler[U, R], U IBaseUseCase, R any](ctx context.Context, handler H, useCase U, result *UseCaseResult[R]) error {
	var (
		handleErr     error
		middlewareErr error
		middleWares   = handler.GetMiddlewares()
		innerResult   *UseCaseResult[R]
	)

	ctx = event_handler.NewEventContext(ctx)

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase = middleWare.Before(ctx, useCase); middlewareErr != nil {
			return middlewareErr
		}
	}

	handleErr, innerResult = handler.Handle(ctx, useCase)

	*result = *innerResult

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase, result = middleWare.After(ctx, useCase, handleErr, result); middlewareErr != nil {
			return middlewareErr
		}
	}

	return handleErr
}
