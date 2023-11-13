package elastic_product_consumer

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/sercanakmaz/go-boilerplate-v3/events/product/products"
	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
)

type ElasticProductConsumer struct {
	MessageBus *rabbitmqv1.Client
	Logger     logger.Logger
	Client     *elasticsearch.Client
}

func NewElasticProductConsumer(messageBus *rabbitmqv1.Client, log logger.Logger, client *elasticsearch.Client) *ElasticProductConsumer {
	return &ElasticProductConsumer{
		MessageBus: messageBus,
		Logger:     log,
		Client:     client,
	}
}

func (c *ElasticProductConsumer) Construct() {
	c.MessageBus.AddConsumer("HG.Elastic.Product.Create").
		SubscriberExchange("*", rabbitmqv1.Topic, "HG.Integration.Product:Created").
		HandleConsumer(c.createElasticProduct())
}

func (c *ElasticProductConsumer) createElasticProduct() func(message rabbitmqv1.Message) error {
	return func(message rabbitmqv1.Message) error {
		var eventMessage products.Created

		ctx := c.Logger.WithCorrelationId(context.Background(), message.GetCorrelationId())

		if err := json.Unmarshal(message.Payload, &eventMessage); err != nil {
			return err
		}

		c.Logger.Info(ctx, "consumer elastic start")

		c.Logger.Info(ctx, eventMessage.Sku)

		data, _ := json.Marshal(eventMessage)

		_, err := c.Client.Index("product_ddd_qa", bytes.NewReader(data))
		if err != nil {
			return err
		}

		c.Logger.Info(ctx, "consumer elastic finish")

		return nil
	}
}
