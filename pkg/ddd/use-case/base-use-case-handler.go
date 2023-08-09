package use_case

import "context"

type IBaseUseCaseHandler[U IBaseUseCase, R any] interface {
	GetMiddlewares() []IBaseUseCaseMiddleware[U, R]
	Handle(ctx context.Context, useCase U) (error, *UseCaseResult[R])
}
