package products

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourMongo "github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IProductRepository interface {
	FindOneById(ctx context.Context, id primitive.ObjectID) (*Product, error)
	FindOneBySku(ctx context.Context, sku string) (*Product, error)
	Delete(ctx context.Context, product *Product) error
	Add(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	UpdateStock(ctx context.Context, product *UpdateProductStock) error
}

const _collectionName = "Products"

type productRepository struct {
	db              *mongo.Database
	eventDispatcher ddd.IEventDispatcher
}

func NewProductRepository(db *mongo.Database) IProductRepository {
	return &productRepository{db: db}
}

func (repository productRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*Product, error) {
	sessionContext := ourMongo.GetSessionContext(ctx)

	var product *Product
	err := repository.db.Collection(_collectionName).FindOne(sessionContext, bson.M{"_id": id}).Decode(&product)
	return product, err
}

func (repository productRepository) FindOneBySku(ctx context.Context, sku string) (*Product, error) {
	sessionContext := ourMongo.GetSessionContext(ctx)

	var product *Product
	err := repository.db.Collection(_collectionName).FindOne(sessionContext, bson.M{"Sku": sku}).Decode(&product)
	return product, err
}

func (repository productRepository) Delete(ctx context.Context, product *Product) error {
	sessionContext := ourMongo.GetSessionContext(ctx)

	_, err := repository.db.Collection(_collectionName).DeleteOne(sessionContext, bson.M{"Sku": product.Sku})

	ddd.DispatchDomainEvents(ctx, product)

	return err
}

func (repository productRepository) Add(ctx context.Context, product *Product) error {
	sessionContext := ourMongo.GetSessionContext(ctx)

	_, err := repository.db.Collection(_collectionName).InsertOne(sessionContext, &product, options.InsertOne())

	ddd.DispatchDomainEvents(ctx, product)

	return err
}

func (repository productRepository) Update(ctx context.Context, product *Product) error {
	sessionContext := ourMongo.GetSessionContext(ctx)

	_, err := repository.db.Collection(_collectionName).ReplaceOne(sessionContext, bson.M{"_id": product.Id}, &product)

	return err
}

func (repository productRepository) UpdateStock(ctx context.Context, product *UpdateProductStock) error {
	sessionContext := ourMongo.GetSessionContext(ctx)

	_, err := repository.db.Collection(_collectionName).UpdateOne(sessionContext, bson.M{"Sku": product.Sku}, bson.M{
		"$set": product,
	})

	return err
}
