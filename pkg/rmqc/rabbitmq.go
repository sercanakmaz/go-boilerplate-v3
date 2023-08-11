package rmqc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	time "time"
)

type RabbitMQ struct {
	config     RabbitMQConfig
	connection *amqp.Connection
	consumers  []*Consumer
	pubChannel *amqp.Channel
}

func newRabbitMQ(config RabbitMQConfig) *RabbitMQ {
	return &RabbitMQ{config: config}
}

func Connect(config RabbitMQConfig) (*RabbitMQ, error) {
	rbt := newRabbitMQ(config)

	connection, err := amqp.DialConfig(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		rbt.config.Username,
		rbt.config.Password,
		rbt.config.Host,
		rbt.config.Port,
		rbt.config.VHost,
	), amqp.Config{Properties: amqp.Table{"connection_name": rbt.config.ConnectionName}})

	if err != nil {
		return nil, err
	}

	rbt.connection = connection

	go rbt.addConnectionCloseListener()

	return rbt, err
}

func (rbt *RabbitMQ) reconnect() error {
	connection, err := amqp.DialConfig(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		rbt.config.Username,
		rbt.config.Password,
		rbt.config.Host,
		rbt.config.Port,
		rbt.config.VHost,
	), amqp.Config{Properties: amqp.Table{"connection_name": rbt.config.ConnectionName}})

	if err != nil {
		return err
	}

	rbt.connection = connection

	go rbt.addConnectionCloseListener()

	return nil
}

func (rbt *RabbitMQ) channel() (*amqp.Channel, error) {
	if rbt.connection != nil && !rbt.connection.IsClosed() {
		return rbt.connection.Channel()
	} else {
		return nil, errors.New("Connection is not opened!")
	}
}

func (rbt *RabbitMQ) channelForPublish() (*amqp.Channel, error) {

	if rbt.pubChannel == nil {
		ch, err := rbt.channel()
		if err != nil {
			return nil, err
		}
		rbt.pubChannel = ch
	}

	return rbt.pubChannel, nil
}

func (rbt *RabbitMQ) BindConsumer(consumerModel IConsumer) error {
	var (
		err      error
		ch       *amqp.Channel
		consumer *Consumer
	)

	cb := NewConsumerBuilder()
	cb.setConsumer(consumerModel)
	consumerModel.Configure(cb)

	if ch, err = rbt.channel(); err != nil {
		return err
	}

	if consumer, err = cb.build(ch); err != nil {
		return err
	}

	rbt.consumers = append(rbt.consumers, consumer)
	return nil
}

func (rbt *RabbitMQ) Start() {
	for _, consumer := range rbt.consumers {
		consumer.channel.Qos(consumer.config.Qos.PrefetchCount, 0, false)
		for i := 0; i < consumer.config.Qos.ConsumerCount; i++ {
			go consumer.Consume(i)
		}
	}
}

func (rbt RabbitMQ) Publish(exchangeName, routingKey string, data interface{}) error {
	ch, err := rbt.channelForPublish()

	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		MessageId:       uuid.New().String(),
		Timestamp:       time.Now(),
		Body:            body,
		ContentType:     "use-cases/json",
		ContentEncoding: "UTF-8",
	})
}

func (rbt RabbitMQ) PublishWithCorrelationId(exchangeName, routingKey, correlationId string, data interface{}) error {
	ch, err := rbt.channelForPublish()

	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		MessageId:       uuid.New().String(),
		CorrelationId:   correlationId,
		Timestamp:       time.Now(),
		Body:            body,
		ContentType:     "use-cases/json",
		ContentEncoding: "UTF-8",
	})
}

func (rbt *RabbitMQ) addConnectionCloseListener() {
	receiver := make(chan *amqp.Error)
	rbt.connection.NotifyClose(receiver)
	err := <-receiver

	if err != nil {
		for i := 0; i < rbt.config.Reconnect.MaxAttempt; i++ {
			if err := rbt.reconnect(); err != nil {
				rbt.Start()
				return
			}

			time.Sleep(rbt.config.Reconnect.Interval)
		}

	}
}
