package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv2"
)

func NewProductRepositoryResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) IProductRepository {
	rbt := rabbitmqv2.NewRabbitMQResolve(rabbitMQConfig)
	eventHandler := rabbitmqv2.NewEventHandlerResolve(rbt)
	db, _ := mongo.NewMongoDb(mongoConfig)
	return newProductRepository(db, eventHandler)
}

func NewProductServiceResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) IProductService {
	return NewProductService(NewProductRepositoryResolve(rabbitMQConfig, mongoConfig))
}
