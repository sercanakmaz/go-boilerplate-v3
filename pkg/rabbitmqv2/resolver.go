package rabbitmqv2

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rmqc"
	"sync"
)

var (
	once = sync.Once{}
	rbt  *rmqc.RabbitMQ
)

func NewEventHandlerResolve(rbt *rmqc.RabbitMQ) ddd.IEventDispatcher {
	return NewRabbitMQEventDispatcher(rbt)
}

func NewRabbitMQResolve(config Config) *rmqc.RabbitMQ {
	var err error

	once.Do(func() {
		rbt, err = rmqc.Connect(rmqc.RabbitMQConfig{
			Host:           config.Host,
			Username:       config.Username,
			Password:       config.Password,
			Port:           config.Port,
			VHost:          config.VHost,
			ConnectionName: config.ConnectionName,
			Reconnect: rmqc.Reconnect{
				MaxAttempt: config.Reconnect.MaxAttempt,
				Interval:   config.Reconnect.Interval,
			},
		})
	})

	if err != nil {
		panic(err)
	}

	return rbt
}
