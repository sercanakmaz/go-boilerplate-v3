package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
)

type RabbitMQEventPublisher struct {
	MessageBus *rabbitmqv1.Client
}

func NewRabbitMQEventPublisher(messageBus *rabbitmqv1.Client) ddd.IEventPublisher {
	return &RabbitMQEventPublisher{MessageBus: messageBus}
}

func (p *RabbitMQEventPublisher) Publish(ctx context.Context) error {
	eventContext := ddd.GetEventContext(ctx)

	for true {
		dispatchedEvent := eventContext.TakeDispatched()
		if dispatchedEvent == nil {
			break
		}

		if err := p.MessageBus.Publish(ctx, "*", dispatchedEvent); err != nil {
			return err
		}
	}

	return nil
}
