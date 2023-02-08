package ddd

import (
	"github.com/rozturac/rmqc"
)

type RabbitMQEventDispatcher struct {
	rbt *rmqc.RabbitMQ
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (handler RabbitMQEventDispatcher) Dispatch(events []IBaseEvent) {
	for _, event := range events {
		handler.rbt.Publish(event.ExchangeName(), "", event)
	}
}
