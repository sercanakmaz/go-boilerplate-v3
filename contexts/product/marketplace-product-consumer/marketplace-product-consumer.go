package marketplace_product_consumer

import (
	"context"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	marketplace_products "github.com/sercanakmaz/go-boilerplate-v3/events/product/marketplace-products"
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
	c.MessageBus.AddConsumer("HG.Marketplace.Product.Create.Publish").
		SubscriberExchange("*", rabbitmqv1.Topic, "HG.Integration.Product:Created").
		HandleConsumer(c.marketplaceProductCreatePublisher())

	// Farklı consumer olabilir.

	c.MessageBus.AddConsumer("HG.Marketplace.Product.Create.Foleja").
		SubscriberExchange("*", rabbitmqv1.Topic, "HG.Integration.MarketplaceProduct:Create").
		HandleConsumer(c.createMarketplaceProductForFoleja())

	c.MessageBus.AddConsumer("HG.Marketplace.Product.Create.Hepsiglobal").
		SubscriberExchange("*", rabbitmqv1.Topic, "HG.Integration.MarketplaceProduct:Create").
		HandleConsumer(c.createMarketplaceProductForHepsiglobal())
}

func (c *MarketplaceProductConsumer) marketplaceProductCreatePublisher() func(message rabbitmqv1.Message) error {
	return func(message rabbitmqv1.Message) error {
		var eventMessage products.Created

		if err := json.Unmarshal(message.Payload, &eventMessage); err != nil {
			return err
		}

		ctx := c.Logger.WithCorrelationId(context.Background(), uuid.NewV4().String())

		c.Logger.Info(ctx, "consumer start")

		// Company'nin aktif olduğu marketplaceleri çektik.

		marketplaces := []string{"Foleja", "Hepsiglobal"}

		for key, _ := range marketplaces {
			if err := c.MessageBus.Publish(ctx, "*", marketplace_products.Create{
				Sku:           eventMessage.Sku,
				CompanyId:     eventMessage.CompanyId,
				MarketplaceId: key,
			}); err != nil {
				return err
			}
		}

		return nil
	}
}

func (c *MarketplaceProductConsumer) createMarketplaceProductForFoleja() func(message rabbitmqv1.Message) error {
	return func(message rabbitmqv1.Message) error {
		var eventMessage marketplace_products.Create

		if err := json.Unmarshal(message.Payload, &eventMessage); err != nil {
			return err
		}

		// Foleja için gelen bir event değilse ignore et.
		if eventMessage.MarketplaceId != 0 {
			return nil
		}

		ctx := c.Logger.WithCorrelationId(context.Background(), uuid.NewV4().String())

		c.Logger.Info(ctx, "consumer start")

		// Foleja Logic

		fmt.Println(eventMessage)

		return nil
	}
}

func (c *MarketplaceProductConsumer) createMarketplaceProductForHepsiglobal() func(message rabbitmqv1.Message) error {
	return func(message rabbitmqv1.Message) error {
		var eventMessage marketplace_products.Create

		if err := json.Unmarshal(message.Payload, &eventMessage); err != nil {
			return err
		}

		// Hepsiglobal için gelen bir event değilse ignore et.
		if eventMessage.MarketplaceId != 1 {
			return nil
		}

		ctx := c.Logger.WithCorrelationId(context.Background(), uuid.NewV4().String())

		c.Logger.Info(ctx, "consumer start")

		// HG Logic

		fmt.Println(eventMessage)

		return nil
	}
}
