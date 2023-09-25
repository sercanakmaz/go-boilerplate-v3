package infra

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
)

type RabbitMQEventPublisher struct {
	MessageBus *rabbitmqv1.Client
	Logger     logger.Logger
}

func NewRabbitMQEventPublisher(messageBus *rabbitmqv1.Client, log logger.Logger) ddd.IEventPublisher {
	return &RabbitMQEventPublisher{MessageBus: messageBus, Logger: log}
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

		p.Logger.InfoWithExtra(ctx, "Published event to RabbitMQ", map[string]interface{}{
			"EventName": dispatchedEvent.EventName(),
			"Payload":   dispatchedEvent,
		})
	}

	return nil
}
