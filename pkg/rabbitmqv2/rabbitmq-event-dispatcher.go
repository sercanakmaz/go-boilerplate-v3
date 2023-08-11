package rabbitmqv2

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/event-handler"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rmqc"
)

type RabbitMQEventDispatcher struct {
	rbt *rmqc.RabbitMQ
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) event_handler.IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (handler RabbitMQEventDispatcher) Dispatch(events []event_handler.IBaseEvent) {
	for _, event := range events {
		handler.rbt.Publish(event.ExchangeName(), "", event)
	}
}
