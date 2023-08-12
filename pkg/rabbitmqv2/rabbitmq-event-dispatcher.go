package rabbitmqv2

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rmqc"
)

type RabbitMQEventDispatcher struct {
	rbt *rmqc.RabbitMQ
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) ddd.IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (dispatcher RabbitMQEventDispatcher) Dispatch(events []ddd.IBaseEvent) {
	for _, event := range events {
		dispatcher.rbt.Publish(event.ExchangeName(), "", event)
	}
}
