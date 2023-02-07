package ddd

import (
	"github.com/rozturac/rmqc"
	"reflect"
)

type RabbitMQEventDispatcher struct {
	rbt *rmqc.RabbitMQ
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (handler RabbitMQEventDispatcher) Dispatch(events []IBaseEvent) {
	for _, event := range events {
		t := reflect.TypeOf(event)
		eventName := t.Elem().Name()
		handler.rbt.Publish(eventName, "", event)
	}
}
