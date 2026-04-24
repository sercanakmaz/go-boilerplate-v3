package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb(config Config) (*mongo.Database, *mongo.Client) {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoDb))

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	return client.Database(config.Database), client
}
