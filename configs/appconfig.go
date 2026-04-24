package configs

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"time"
)

type Config struct {
	Mongo         mongo.Config
	Host          Host
	RabbitMQ      RabbitMQConfig
	ElasticSearch ElasticSearch
}

type Host struct {
	Port int
}

type RabbitMQConfig struct {
	Host           string
	Port           string
	VHost          string
	Username       string
	Password       string
	ConnectionName string
	Reconnect      struct {
		MaxAttempt int
		Interval   time.Duration
	}
}

type ElasticSearch struct {
	Host string
	Port string
}
