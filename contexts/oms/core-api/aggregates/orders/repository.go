package orders

import (
	"context"
	event_handler "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOrderRepository interface {
	FindOneById(ctx context.Context, id primitive.ObjectID) (*Order, error)
	FindOneByOrderNumber(ctx context.Context, orderNumber string) (*Order, error)
	Add(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
}

const _collectionName = "Orders"

type orderRepository struct {
	db *mongo.Database
}

func NewOrderRepository(db *mongo.Database) IOrderRepository {
	return &orderRepository{db: db}
}

func (repository orderRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*Order, error) {
	var order *Order
	err := repository.db.Collection(_collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	return order, err
}

func (repository orderRepository) FindOneByOrderNumber(ctx context.Context, orderNumber string) (*Order, error) {
	var order *Order
	err := repository.db.Collection(_collectionName).FindOne(ctx, bson.M{"OrderNumber": orderNumber}).Decode(&order)
	return order, err
}

func (repository orderRepository) Add(ctx context.Context, order *Order) error {
	if _, err := repository.db.Collection(_collectionName).InsertOne(ctx, &order, options.InsertOne()); err != nil {
		return err
	}

	event_handler.DispatchDomainEvents(ctx, order)

	return nil
}

func (repository orderRepository) Update(ctx context.Context, order *Order) error {
	if _, err := repository.db.Collection(_collectionName).ReplaceOne(ctx, bson.M{"_id": order.Id}, &order); err != nil {
		return err
	}

	event_handler.DispatchDomainEvents(ctx, order)

	return nil
}
