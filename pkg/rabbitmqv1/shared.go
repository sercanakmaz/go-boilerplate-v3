package rabbitmqv1

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func getGuid() string {
	return uuid.New().String()
}

func isGuid(p string) bool {
	_, err := uuid.Parse(p)
	if err != nil {
		return false
	}
	return true
}

func convertExchangeType(exchangeType ExchangeType) string {

	var rabbitmqExchangeType string
	switch exchangeType {
	case Direct:
		rabbitmqExchangeType = amqp.ExchangeDirect
		break
	case Fanout:
		rabbitmqExchangeType = amqp.ExchangeFanout
		break
	case Topic:
		rabbitmqExchangeType = amqp.ExchangeTopic
		break
	case ConsistentHashing:
		rabbitmqExchangeType = "x-consistent-hash"
		break
	case XDelayedMessage:
		rabbitmqExchangeType = "x-delayed-message"
		break
	}
	return rabbitmqExchangeType
}
