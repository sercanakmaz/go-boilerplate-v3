package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDb(config Config) (*mongo.Database, *mongo.Client) {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoDb))

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	return client.Database(config.Database), client
}
