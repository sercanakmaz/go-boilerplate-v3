package created_log_consumer

import (
	"github.com/sercanakmaz/go-boilerplate-v3/configs"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/config"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv2"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	var (
		cfg configs.Config
		err error
	)

	if err = config.Load(&cfg); err != nil {
		return err
	}
	rbt := rabbitmqv2.NewRabbitMQResolve(cfg.RabbitMQ)

	if err = rbt.BindConsumer(NewUserCreatedConsumer()); err != nil {
		return err
	}

	rbt.Start()

	return nil
}
