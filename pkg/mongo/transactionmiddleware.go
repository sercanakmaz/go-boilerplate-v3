package mongo

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"go.mongodb.org/mongo-driver/mongo"
)

var sessionContextKey = "mongoSessionContext"

func GetSessionContext(ctx context.Context) mongo.SessionContext {
	return ctx.Value(sessionContextKey).(mongo.SessionContext)
}

type TransactionMiddleware[T ddd.IBaseUseCase, R any] struct {
	client  *mongo.Client
	session mongo.Session
	sc      mongo.SessionContext
}

func NewTransactionMiddleware[T ddd.IBaseUseCase, R any](client *mongo.Client) *TransactionMiddleware[T, R] {
	return &TransactionMiddleware[T, R]{client: client}
}

func (s *TransactionMiddleware[T, R]) Before(ctx context.Context, useCase T) (error, context.Context, T) {
	var (
		err error
	)

	if s.session, err = s.client.StartSession(); err != nil {
		return err, ctx, useCase
	}

	if err = s.session.StartTransaction(); err != nil {
		return err, ctx, useCase
	}

	s.sc = mongo.NewSessionContext(ctx, s.session)

	ctx = context.WithValue(ctx, sessionContextKey, s.sc)

	return nil, ctx, useCase
}

func (s *TransactionMiddleware[T, R]) After(ctx context.Context, useCase T, err error, result *ddd.UseCaseResult[R]) (error, context.Context, T, *ddd.UseCaseResult[R]) {

	if err != nil {
		if err = s.session.AbortTransaction(s.sc); err != nil {
			return err, ctx, useCase, result
		}
	} else {
		if err = s.session.CommitTransaction(s.sc); err != nil {
			return err, ctx, useCase, result
		}
	}

	s.session.EndSession(ctx)

	return nil, ctx, useCase, result
}
