package orders

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
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
	db              *mongo.Database
	eventDispatcher ddd.IEventDispatcher
}

func newOrderRepository(db *mongo.Database, eventDispatcher ddd.IEventDispatcher) IOrderRepository {
	return &orderRepository{db: db, eventDispatcher: eventDispatcher}
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
	_, err := repository.db.Collection(_collectionName).InsertOne(ctx, &order, options.InsertOne())
	repository.dispatchDomainEvents(order, err)

	return err
}

func (repository orderRepository) Update(ctx context.Context, order *Order) error {
	_, err := repository.db.Collection(_collectionName).ReplaceOne(ctx, bson.M{"_id": order.Id}, &order)
	repository.dispatchDomainEvents(order, err)

	return err
}

func (repository orderRepository) dispatchDomainEvents(order *Order, err error) {
	if err == nil {
		repository.eventDispatcher.Dispatch(order.GetDomainEvents())
		order.ClearDomainEvents()
	}
}
