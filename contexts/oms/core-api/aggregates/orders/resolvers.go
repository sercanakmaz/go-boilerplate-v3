package orders

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv2"
)

func NewOrderRepositoryResolve(mongoConfig mongo.Config) IOrderRepository {
	return newOrderRepository(mongo.NewMongoDb(mongoConfig))
}

func NewOrderServiceResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) IOrderService {
	return NewOrderService(NewOrderRepositoryResolve(mongoConfig))
}
