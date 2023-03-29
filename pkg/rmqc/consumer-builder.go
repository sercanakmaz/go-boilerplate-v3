package rmqc

import (
	"github.com/streadway/amqp"
	"time"
)

type ConsumerBuilder struct {
	consumerName  string
	consumerCount int
	prefetchCount int
	queue         *Queue
	subscribers   []*Subscriber
	consumer      IConsumer
	Reconnect     struct {
		MaxAttempt int
		Interval   time.Duration
	}
}

func NewConsumerBuilder() *ConsumerBuilder {
	return &ConsumerBuilder{
		prefetchCount: 1,
		consumerCount: 1,
		Reconnect: Reconnect{
			MaxAttempt: 1,
			Interval:   1,
		},
	}
}

func (cb *ConsumerBuilder) BindQueue(name string) *ConsumerBuilder {
	cb.BindQueueWithArg(name, nil)
	return cb
}

func (cb *ConsumerBuilder) BindQueueWithArg(name string, arg map[string]interface{}) *ConsumerBuilder {
	cb.queue = newQueue(name, arg)
	return cb
}

func (cb *ConsumerBuilder) Subscribe(exchangeName, routingKey string, exchangeType ExchangeType) *ConsumerBuilder {
	subscriber := newSubscriber(exchangeName, routingKey, exchangeType)
	cb.subscribers = append(cb.subscribers, subscriber)
	return cb
}

func (cb *ConsumerBuilder) SubscribeAsTopic(exchangeName, routingKey string) *ConsumerBuilder {
	subscriber := newSubscriber(exchangeName, routingKey, ExchangeType_Topic)
	cb.subscribers = append(cb.subscribers, subscriber)
	return cb
}

func (cb *ConsumerBuilder) SubscribeAsDirect(exchangeName, routingKey string) *ConsumerBuilder {
	subscriber := newSubscriber(exchangeName, routingKey, ExchangeType_Direct)
	cb.subscribers = append(cb.subscribers, subscriber)
	return cb
}

func (cb *ConsumerBuilder) SubscribeAsFanout(exchangeName, routingKey string) *ConsumerBuilder {
	subscriber := newSubscriber(exchangeName, routingKey, ExchangeType_Fanout)
	cb.subscribers = append(cb.subscribers, subscriber)
	return cb
}

func (cb *ConsumerBuilder) SubscribeAsHeaders(exchangeName, routingKey string) *ConsumerBuilder {
	subscriber := newSubscriber(exchangeName, routingKey, ExchangeType_Headers)
	cb.subscribers = append(cb.subscribers, subscriber)
	return cb
}

func (cb *ConsumerBuilder) setConsumer(consumer IConsumer) *ConsumerBuilder {
	cb.consumer = consumer
	return cb
}

func (cb *ConsumerBuilder) SetConsumerName(v string) *ConsumerBuilder {
	cb.consumerName = v
	return cb
}

func (cb *ConsumerBuilder) SetConsumerCount(v int) *ConsumerBuilder {
	cb.consumerCount = v
	return cb
}

func (cb *ConsumerBuilder) SetPrefetchCount(v int) *ConsumerBuilder {
	cb.prefetchCount = v
	return cb
}

func (cb *ConsumerBuilder) build(ch *amqp.Channel) (consumer *Consumer, err error) {
	if err = cb.queue.declare(ch); err != nil {
		return
	}

	for _, subscriber := range cb.subscribers {
		if err = subscriber.declare(ch); err != nil {
			return
		}

		if err = cb.queue.exchangeBind(ch, subscriber); err != nil {
			return
		}
	}

	consumer = newConsumer(ConsumerConfig{
		QueueName: cb.queue.Name,
		Consumer:  cb.consumer,
		Qos: struct {
			PrefetchCount int
			ConsumerCount int
		}{PrefetchCount: cb.prefetchCount, ConsumerCount: cb.consumerCount},
	}, ch)

	return
}
