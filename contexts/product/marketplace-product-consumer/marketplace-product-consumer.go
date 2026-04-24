package marketplace_product_consumer

import (
	"context"

	"github.com/sercanakmaz/go-boilerplate-v3/events/product/products"
	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
)

type MarketplaceProductConsumer struct {
	MessageBus *rabbitmqv1.Client
	Logger     logger.Logger
}

func NewMarketplaceProductConsumer(messageBus *rabbitmqv1.Client, log logger.Logger) *MarketplaceProductConsumer {
	return &MarketplaceProductConsumer{
		MessageBus: messageBus,
		Logger:     log,
	}
}

func (c *MarketplaceProductConsumer) Construct() {
	c.MessageBus.AddConsumer("HG.Marketplace.Product.Create.Foleja").
		SubscriberExchange("*", rabbitmqv1.Topic, "HG.Integration.Product:Created").
		HandleConsumer(c.createMarketplaceProductForFoleja())

	c.MessageBus.AddConsumer("HG.Marketplace.Product.Create.Hepsiglobal").
		SubscriberExchange("*", rabbitmqv1.Topic, "HG.Integration.Product:Created").
		HandleConsumer(c.createMarketplaceProductForHepsiglobal())
}

func (c *MarketplaceProductConsumer) createMarketplaceProductForFoleja() func(message rabbitmqv1.Message) error {
	return func(message rabbitmqv1.Message) error {
		ctx := c.Logger.WithCorrelationId(context.Background(), message.GetCorrelationId())

		eventMessage, err := rabbitmqv1.DecodeMessage[products.Created](message)
		if err != nil {
			return err
		}

		// Company'nin aktif olduğu marketplaceleri çektik.

		// MP'ler içinde Foleja yoksa eventi ignore et.

		c.Logger.Info(ctx, "consumer start")
		c.Logger.Info(ctx, eventMessage.Sku)

		// Foleja Logic

		c.Logger.Info(ctx, "consumer finish")

		return nil
	}
}

func (c *MarketplaceProductConsumer) createMarketplaceProductForHepsiglobal() func(message rabbitmqv1.Message) error {
	return func(message rabbitmqv1.Message) error {
		ctx := c.Logger.WithCorrelationId(context.Background(), message.GetCorrelationId())

		eventMessage, err := rabbitmqv1.DecodeMessage[products.Created](message)
		if err != nil {
			return err
		}

		// Company'nin aktif olduğu marketplaceleri çektik.

		// MP'ler içinde HG yoksa eventi ignore et.

		c.Logger.Info(ctx, "consumer start")
		c.Logger.Info(ctx, eventMessage.Sku)

		// HG Logic

		c.Logger.Info(ctx, "consumer finish")
		return nil
	}
}
