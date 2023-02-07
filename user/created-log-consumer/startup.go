package created_log_consumer

import (
	"go-boilerplate-v3/configs"
	"go-boilerplate-v3/pkg/config"
	"go-boilerplate-v3/pkg/rabbitmqv2"
	"runtime"
)

func Init() {
	var (
		cfg configs.Config
		err error
	)

	if err = config.Load(&cfg); err != nil {
		panic(err)
	}

	rbt := rabbitmqv2.NewRabbitMQResolve(cfg.RabbitMQ)
	if err := rbt.BindConsumer(NewUserCreatedConsumer()); err != nil {
		panic(err)
	}
	rbt.Start()
	runtime.Goexit()
}
