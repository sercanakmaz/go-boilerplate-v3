package configs

import (
	"go-boilerplate-v3/pkg/mongo"
	"go-boilerplate-v3/pkg/rabbitmqv2"
)

type Config struct {
	Mongo    mongo.Config
	Host     Host
	RabbitMQ rabbitmqv2.Config
}

type Host struct {
	Port int
}
