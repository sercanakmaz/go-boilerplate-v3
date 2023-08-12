package mongo

import (
	"context"
	"fmt"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type TransactionMiddleware[T ddd.IBaseUseCase, R any] struct {
}

func (self *TransactionMiddleware[T, R]) Before(ctx context.Context, useCase T) (error, context.Context, T) {
	ctx = context.WithValue(ctx, "transaction", "mongostart")
	return nil, ctx, useCase
}
func (self *TransactionMiddleware[T, R]) After(ctx context.Context, useCase T, err error, result *ddd.UseCaseResult[R]) (error, context.Context, T, *ddd.UseCaseResult[R]) {
	fmt.Println(ctx.Value("transaction"))
	ctx = context.WithValue(ctx, "transaction", "mongoend")
	return nil, ctx, useCase, result
}
