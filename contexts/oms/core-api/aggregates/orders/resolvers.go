package orders

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv2"
)

func NewOrderRepositoryResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) IOrderRepository {
	rbt := rabbitmqv2.NewRabbitMQResolve(rabbitMQConfig)
	eventHandler := rabbitmqv2.NewEventHandlerResolve(rbt)
	return newOrderRepository(mongo.NewMongoDb(mongoConfig), eventHandler)
}

func NewOrderServiceResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) IOrderService {
	return NewOrderService(NewOrderRepositoryResolve(rabbitMQConfig, mongoConfig))
}
