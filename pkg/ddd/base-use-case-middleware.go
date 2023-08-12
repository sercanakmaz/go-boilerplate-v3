package ddd

import "context"

type IBaseUseCaseMiddleware[T IBaseUseCase, R any] interface {
	Before(ctx context.Context, useCase T) (error, context.Context, T)
	After(ctx context.Context, useCase T, err error, result *UseCaseResult[R]) (error, context.Context, T, *UseCaseResult[R])
}
