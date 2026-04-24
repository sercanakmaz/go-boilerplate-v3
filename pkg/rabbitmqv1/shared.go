package rabbitmqv1

import (
	"encoding/json"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
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

func DecodeMessage[T any](m Message) (T, error) {
	var result T
	err := json.Unmarshal(m.Payload, &result)
	return result, err
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
