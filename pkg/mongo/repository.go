package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionNamer interface {
	CollectionName() string
}

// MongoRepository[T] is a generic base repository that handles common CRUD
// operations for any entity type T against a single MongoDB collection.
// Embed it in domain-specific repository structs and call its methods
// instead of repeating the same Collection/Decode pattern everywhere.
type MongoRepository[T any] struct {
	DB             *mongo.Database
	CollectionName string
}

func NewMongoRepository[T any](db *mongo.Database) *MongoRepository[T] {
	return &MongoRepository[T]{
		DB:             db,
		CollectionName: collectionNameFromType[T](),
	}
}

func (r *MongoRepository[T]) FindOneByFilter(ctx context.Context, filter bson.M) (*T, error) {
	sc := GetSessionContext(ctx)
	var result T
	err := r.DB.Collection(r.CollectionName).FindOne(sc, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *MongoRepository[T]) InsertOne(ctx context.Context, entity *T) error {
	sc := GetSessionContext(ctx)
	_, err := r.DB.Collection(r.CollectionName).InsertOne(sc, entity, options.InsertOne())
	return err
}

func (r *MongoRepository[T]) DeleteOne(ctx context.Context, filter bson.M) error {
	sc := GetSessionContext(ctx)
	_, err := r.DB.Collection(r.CollectionName).DeleteOne(sc, filter)
	return err
}

func (r *MongoRepository[T]) ReplaceOne(ctx context.Context, filter bson.M, entity *T) error {
	sc := GetSessionContext(ctx)
	_, err := r.DB.Collection(r.CollectionName).ReplaceOne(sc, filter, entity)
	return err
}

func (r *MongoRepository[T]) UpdateOne(ctx context.Context, filter bson.M, update bson.M) error {
	sc := GetSessionContext(ctx)
	_, err := r.DB.Collection(r.CollectionName).UpdateOne(sc, filter, update)
	return err
}
