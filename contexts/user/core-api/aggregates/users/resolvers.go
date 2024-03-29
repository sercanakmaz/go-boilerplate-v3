package users

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv2"
)

func NewUserRepositoryResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) IUserRepository {
	rbt := rabbitmqv2.NewRabbitMQResolve(rabbitMQConfig)
	eventHandler := rabbitmqv2.NewEventHandlerResolve(rbt)
	db, _ := mongo.NewMongoDb(mongoConfig)

	return newUserRepository(db, eventHandler)
}

func NewUserServiceResolve(rabbitMQConfig rabbitmqv2.Config, mongoConfig mongo.Config) UserService {
	return NewUserService(NewUserRepositoryResolve(rabbitMQConfig, mongoConfig))
}
