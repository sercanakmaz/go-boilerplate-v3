package marketplace_product_consumer

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/configs"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/docs"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/config"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	docs.Initialize()

	ctx := context.Background()

	var (
		cfg configs.Config
		err error
	)

	if err = config.Load(&cfg); err != nil {
		return err
	}

	var logger = log.NewLogger()

	messageBus := rabbitmqv1.NewRabbitMqClient(
		[]string{cfg.RabbitMQ.Host},
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		"",
		logger,
		rabbitmqv1.RetryCount(0),
		rabbitmqv1.PrefetchCount(1))

	consumer := NewMarketplaceProductConsumer(messageBus, logger)
	consumer.Construct()

	return messageBus.RunConsumers(ctx)
}
