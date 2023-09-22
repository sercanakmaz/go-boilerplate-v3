package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/events/product/products"
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

		// TODO: gerek kalmadı, kaldırılabilir burası
		if dispatchedEvent.EventName() == "Product:Created" {
			castedEvent := dispatchedEvent.(*products.Created)
			if err := p.MessageBus.Publish(ctx, "*", *castedEvent); err != nil {
				return err
			}
		}
	}

	return nil
}
