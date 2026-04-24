package products

import (
	"context"

	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourMongo "github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	FindOneById(ctx context.Context, id primitive.ObjectID) (*Product, error)
	FindOneBySku(ctx context.Context, sku string) (*Product, error)
	Delete(ctx context.Context, product *Product) error
	Add(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	UpdateStock(ctx context.Context, product *Product) error
}

type productRepository struct {
	*ourMongo.MongoRepository[Product]
	eventDispatcher ddd.IEventDispatcher
}

func NewProductRepository(db *mongo.Database) IProductRepository {
	return &productRepository{
		MongoRepository: ourMongo.NewMongoRepository[Product](db),
	}
}

func (repository productRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*Product, error) {
	return repository.FindOneByFilter(ctx, bson.M{"_id": id})
}

func (repository productRepository) FindOneBySku(ctx context.Context, sku string) (*Product, error) {
	return repository.FindOneByFilter(ctx, bson.M{"Sku": sku})
}

func (repository productRepository) Delete(ctx context.Context, product *Product) error {
	err := repository.DeleteOne(ctx, bson.M{"Sku": product.Sku})

	ddd.DispatchDomainEvents(ctx, product)

	return err
}

func (repository productRepository) Add(ctx context.Context, product *Product) error {
	err := repository.InsertOne(ctx, product)

	ddd.DispatchDomainEvents(ctx, product)

	return err
}

func (repository productRepository) Update(ctx context.Context, product *Product) error {
	return repository.ReplaceOne(ctx, bson.M{"_id": product.Id}, product)
}

func (repository productRepository) UpdateStock(ctx context.Context, product *Product) error {
	err := repository.UpdateOne(ctx, bson.M{"Sku": product.Sku}, bson.M{"$set": bson.M{"stock": product.Stock}})

	ddd.DispatchDomainEvents(ctx, product)

	return err
}
