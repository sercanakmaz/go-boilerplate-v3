package rabbitmqv2

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rmqc"
)

type RabbitMQEventDispatcher struct {
	rbt *rmqc.RabbitMQ
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) ddd.IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (dispatcher RabbitMQEventDispatcher) Dispatch(ctx context.Context, event ddd.IBaseEvent) error {
	return dispatcher.rbt.Publish(event.EventName(), "", event)
}
