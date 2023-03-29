package rmqc

import (
	"github.com/streadway/amqp"
)

type Subscriber struct {
	exchangeName string
	routingKey   string
	exchangeType ExchangeType
}

func newSubscriber(exchangeName, routingKey string, exchangeType ExchangeType) *Subscriber {
	return &Subscriber{exchangeName: exchangeName, routingKey: routingKey, exchangeType: exchangeType}
}

func (s *Subscriber) declare(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(s.exchangeName,
		string(s.exchangeType),
		true,
		false,
		false,
		false,
		nil)
}
