package ddd

import "context"

type IBaseUseCaseHandler[U IBaseUseCase, R any] interface {
	SetMiddlewares()
	GetMiddlewares() []IBaseUseCaseMiddleware[U, R]
	Handle(ctx context.Context, useCase U) (error, *UseCaseResult[R])
}

type BaseUseCaseHandler[U IBaseUseCase, R any] struct {
	middlewares []IBaseUseCaseMiddleware[U, R]
}

func (h *BaseUseCaseHandler[U, R]) Use(middlewares ...IBaseUseCaseMiddleware[U, R]) {
	h.middlewares = append(h.middlewares, middlewares...)
}

func (h *BaseUseCaseHandler[U, R]) GetMiddlewares() []IBaseUseCaseMiddleware[U, R] {
	return h.middlewares
}
