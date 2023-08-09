package mongo

import (
	"context"
	"fmt"
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/use-case"
)

type TransactionMiddleware[T use_case.IBaseUseCase, R any] struct {
}

func (self *TransactionMiddleware[T, R]) Before(ctx context.Context, useCase T) (error, context.Context, T) {
	ctx = context.WithValue(ctx, "transaction", "mongostart")
	return nil, ctx, useCase
}
func (self *TransactionMiddleware[T, R]) After(ctx context.Context, useCase T, err error, result *use_case.UseCaseResult[R]) (error, context.Context, T, *use_case.UseCaseResult[R]) {
	fmt.Println(ctx.Value("transaction"))
	ctx = context.WithValue(ctx, "transaction", "mongoend")
	return nil, ctx, useCase, result
}
