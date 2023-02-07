package rabbitmqv1

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go-boilerplate-v3/pkg/log"
	"golang.org/x/net/context"
	"time"
)

type (
	MessageBroker interface {
		CreateChannel() (*Channel, error)
		CreateConnection(ctx context.Context, parameters MessageBrokerParameter) error
		SignalConnectionStatus(status bool)
		SignalConnection() chan bool
		IsConnected() bool
	}

	MessageBrokerParameter struct {
		Nodes             []string
		PrefetchCount     int
		RetryCount        int
		RetryInterval     time.Duration
		ConcurrentLimit   int
		UserName          string
		Password          string
		selectedNodeIndex int
		VirtualHost       string
	}

	Channel struct {
		rabbitChannel   *amqp.Channel
		prefetchCount   int
		retryCount      int
		concurrentLimit int
		notifyConfirm   chan amqp.Confirmation
	}
	broker struct {
		parameters        MessageBrokerParameter
		connection        *amqp.Connection
		connNotifyChannel chan bool
		logger            log.Logger
	}
)

func (b *broker) CreateConnection(ctx context.Context, parameters MessageBrokerParameter) error {

	var err error

	b.parameters = parameters

	for {

		if b.connection, err = amqp.Dial(b.chooseNode(ctx)); err != nil {
			time.Sleep(reconnectDelay)
			b.logger.Exception(ctx, "Application Retried To Connect RabbitMq", err)
			continue
		}
		b.onClose(ctx)
		b.logger.Info(ctx, "Application  Connected RabbitMq")
		b.SignalConnectionStatus(true)

		break
	}

	return err

}

func (b *broker) chooseNode(ctx context.Context) string {

	if b.parameters.selectedNodeIndex == len(b.parameters.Nodes) {
		b.parameters.selectedNodeIndex = 0
	}

	var selectedNode = b.parameters.Nodes[b.parameters.selectedNodeIndex]
	b.parameters.selectedNodeIndex++
	b.logger.Info(ctx, fmt.Sprintf("Started To Listen On Node %s", selectedNode))
	return fmt.Sprintf("amqp://%s:%s@%s/%s", b.parameters.UserName, b.parameters.Password, selectedNode, b.parameters.VirtualHost)

}

func (b *broker) onClose(ctx context.Context) {
	go func() {
		err := <-b.connection.NotifyClose(make(chan *amqp.Error))
		if err != nil {
			b.logger.Exception(ctx, "Rabbitmq Connection is Down", err)
			b.SignalConnectionStatus(false)
			return
		}
	}()
}

func (b *broker) IsConnected() bool {
	if b.connection == nil {
		return false
	}
	return b.connection.IsClosed() == false
}

func (b *broker) CreateChannel() (*Channel, error) {

	rabbitChannel, err := b.connection.Channel()
	if err != nil {
		return nil, err
	}

	var brokerChannel = Channel{
		rabbitChannel:   rabbitChannel,
		prefetchCount:   b.parameters.PrefetchCount,
		retryCount:      b.parameters.RetryCount,
		concurrentLimit: b.parameters.ConcurrentLimit,
		notifyConfirm:   make(chan amqp.Confirmation, 1),
	}

	return &brokerChannel, nil
}

func (b *broker) SignalConnectionStatus(status bool) {
	go func() {
		b.connNotifyChannel <- status
	}()
}

func (b *broker) SignalConnection() chan bool {
	return b.connNotifyChannel
}

func NewMessageBroker(logger log.Logger) MessageBroker {
	brokerClient := &broker{connNotifyChannel: make(chan bool), logger: logger}
	brokerClient.SignalConnectionStatus(false)

	return brokerClient
}

func (c *Channel) createExchange(exchange string, exchangeType ExchangeType, args amqp.Table) *Channel {
	c.rabbitChannel.ExchangeDeclare(exchange, convertExchangeType(exchangeType), true, false, false, false, args)
	return c

}

func (c *Channel) createQueue(queueName string) *Channel {
	c.rabbitChannel.QueueDeclare(queueName, true, false, false, false, nil)
	return c

}

func (c *Channel) exchangeToQueueBind(exchange string, queueName string, routingKey string, exchangeType ExchangeType) *Channel {
	c.rabbitChannel.ExchangeDeclare(exchange, convertExchangeType(exchangeType), true, false, false, false, nil)
	c.rabbitChannel.QueueBind(queueName, routingKey, exchange, false, nil)
	return c
}

func (c *Channel) exchangeToConsistentExchangeBind(destinationExchange string, sourceExchange string, routingKey string, exchangeType ExchangeType) *Channel {
	c.rabbitChannel.ExchangeDeclare(sourceExchange, convertExchangeType(exchangeType), true, false, false, false, nil)
	c.rabbitChannel.ExchangeBind(destinationExchange, routingKey, sourceExchange, false, nil)
	return c
}
func (c *Channel) createErrorQueueAndBind(errorExchangeName string, errorQueueName string) *Channel {
	c.rabbitChannel.ExchangeDeclare(errorExchangeName, "fanout", true, false, false, false, nil)
	q, _ := c.rabbitChannel.QueueDeclare(errorQueueName, true, false, false, false, nil)
	c.rabbitChannel.QueueBind(q.Name, "", errorExchangeName, false, nil)
	return c

}

func (p *Channel) Publish(ctx context.Context, routingKey string, exchangeName string, payload interface{}) error {
	var message, err = publishExchangeToMessage("", payload)

	if err != nil {
		return err
	}
	return p.rabbitChannel.Publish(exchangeName, routingKey, false, false, message)
}

func publishExchangeToMessage(correlationId string, payload interface{}) (amqp.Publishing, error) {

	var (
		bytes []byte
		err   error
	)

	if bytes, err = json.Marshal(payload); err != nil {
		return amqp.Publishing{}, err
	}

	headers := make(map[string]interface{})
	headers[headerTime] = time.Now().String()
	return amqp.Publishing{
		MessageId:       getGuid(),
		Body:            bytes,
		Headers:         headers,
		CorrelationId:   correlationId,
		Timestamp:       time.Now(),
		DeliveryMode:    deliveryMode,
		ContentEncoding: "UTF-8",
		ContentType:     "application/json",
	}, nil
}

func (b Channel) listenToQueue(queueName string) (<-chan amqp.Delivery, error) {

	msg, err := b.rabbitChannel.Consume(queueName,
		"",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		return nil, err
	}

	return msg, nil
}
