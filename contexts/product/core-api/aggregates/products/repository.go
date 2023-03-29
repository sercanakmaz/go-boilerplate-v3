package products

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IProductRepository interface {
	FindOneById(ctx context.Context, id primitive.ObjectID) (*Product, error)
	FindOneBySku(ctx context.Context, sku string) (*Product, error)
	Add(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
}

const _collectionName = "Products"

type productRepository struct {
	db              *mongo.Database
	eventDispatcher ddd.IEventDispatcher
}

func newProductRepository(db *mongo.Database, eventDispatcher ddd.IEventDispatcher) IProductRepository {
	return &productRepository{db: db, eventDispatcher: eventDispatcher}
}

func (repository productRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*Product, error) {
	var product *Product
	err := repository.db.Collection(_collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	return product, err
}

func (repository productRepository) FindOneBySku(ctx context.Context, sku string) (*Product, error) {
	var product *Product
	err := repository.db.Collection(_collectionName).FindOne(ctx, bson.M{"Sku": sku}).Decode(&product)
	return product, err
}

func (repository productRepository) Add(ctx context.Context, product *Product) error {
	_, err := repository.db.Collection(_collectionName).InsertOne(ctx, &product, options.InsertOne())
	repository.dispatchDomainEvents(product, err)

	return err
}

func (repository productRepository) Update(ctx context.Context, product *Product) error {
	_, err := repository.db.Collection(_collectionName).ReplaceOne(ctx, bson.M{"_id": product.Id}, &product)
	repository.dispatchDomainEvents(product, err)

	return err
}

func (repository productRepository) dispatchDomainEvents(product *Product, err error) {
	if err == nil {
		repository.eventDispatcher.Dispatch(product.GetDomainEvents())
		product.ClearDomainEvents()
	}
}
