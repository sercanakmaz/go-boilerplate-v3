package orderlines

import (
	"context"
	event_handler "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourMongo "github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOrderLineRepository interface {
	FindOneById(ctx context.Context, id primitive.ObjectID) (*OrderLine, error)
	FindByOrderNumber(ctx context.Context, orderNumber string) ([]*OrderLine, error)
	Add(ctx context.Context, orderLine *OrderLine) error
	Update(ctx context.Context, orderLine *OrderLine) error
}

const _collectionName = "OrderLines"

type orderLineRepository struct {
	db *mongo.Database
}

func NewOrderLineRepository(db *mongo.Database) IOrderLineRepository {
	return &orderLineRepository{db: db}
}

func (repository orderLineRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*OrderLine, error) {
	var orderLine *OrderLine

	sessionContext := ourMongo.GetSessionContext(ctx)

	err := repository.db.Collection(_collectionName).FindOne(sessionContext, bson.M{"_id": id}).Decode(&orderLine)

	return orderLine, err
}

func (repository orderLineRepository) FindByOrderNumber(ctx context.Context, orderNumber string) ([]*OrderLine, error) {
	var (
		err        error
		cur        *mongo.Cursor
		orderLines []*OrderLine
	)

	sessionContext := ourMongo.GetSessionContext(ctx)

	if cur, err = repository.db.Collection(_collectionName).Find(sessionContext, bson.M{"OrderNumber": orderNumber}); err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &orderLines); err != nil {
		return nil, err
	}

	return orderLines, err
}

func (repository orderLineRepository) Add(ctx context.Context, orderLine *OrderLine) error {
	sessionContext := ourMongo.GetSessionContext(ctx)

	if _, err := repository.db.Collection(_collectionName).InsertOne(sessionContext, &orderLine, options.InsertOne()); err != nil {
		return err
	}

	event_handler.DispatchDomainEvents(ctx, orderLine)

	return nil
}

func (repository orderLineRepository) Update(ctx context.Context, orderLine *OrderLine) error {
	sessionContext := ourMongo.GetSessionContext(ctx)

	if _, err := repository.db.Collection(_collectionName).ReplaceOne(sessionContext, bson.M{"_id": orderLine.Id}, &orderLine); err != nil {
		return err
	}

	event_handler.DispatchDomainEvents(ctx, orderLine)

	return nil
}
