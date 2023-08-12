package ddd

import (
	"context"
)

func HandleUseCase[H IBaseUseCaseHandler[U, R], U IBaseUseCase, R any](ctx context.Context, handler H, useCase U, result *UseCaseResult[R]) error {
	var (
		handleErr     error
		middlewareErr error
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

	for _, middleWare := range middleWares {
		if middlewareErr, ctx, useCase, result = middleWare.After(ctx, useCase, handleErr, result); middlewareErr != nil {
			return middlewareErr
		}
	}

	return handleErr
}
