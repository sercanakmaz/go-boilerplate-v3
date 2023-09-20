package rabbitmqv1

import (
	"fmt"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"reflect"
	"runtime"
	"time"
)

var (
	ERRORPREFIX     = "_error"
	CONCURRENTLIMIT = 1
	RETRYCOUNT      = 0
	PREFECTCOUNT    = 1
	reconnectDelay  = 2 * time.Second
)

const (
	Direct            ExchangeType = 1
	Fanout            ExchangeType = 2
	Topic             ExchangeType = 3
	ConsistentHashing ExchangeType = 4
	XDelayedMessage   ExchangeType = 5
)

type (
	ExchangeType int

	Client struct {
		context            context.Context
		shutdownFn         context.CancelFunc
		childRoutines      *errgroup.Group
		parameters         MessageBrokerParameter
		shutdownReason     string
		shutdownInProgress bool
		consumerBuilder    ConsumerBuilder
		publisherBuilder   *PublisherBuilder
	}

	withFunc func(*Client) error
)

func (c *Client) checkConsumerConnection(ctx context.Context) {

	go func() {

		for {
			select {

			case isConnected := <-c.consumerBuilder.messageBroker.SignalConnection():
				if !isConnected {
					c.consumerBuilder.messageBroker.CreateConnection(ctx, c.parameters)
					for _, consumer := range c.consumerBuilder.consumers {
						consumer.startConsumerCn <- true
					}
				}
			}
		}

	}()
}

func (c *Client) checkPublisherConnection(ctx context.Context) {

	if c.publisherBuilder.isAlreadyStartConnection {
		return
	}

	c.publisherBuilder.isAlreadyStartConnection = true

	go func() {

		for {

			select {

			case isConnected := <-c.publisherBuilder.messageBroker.SignalConnection():

				if !isConnected {
					c.publisherBuilder.messageBroker.CreateConnection(ctx, c.parameters)
					c.publisherBuilder.startPublisherCh <- true

				}
			}
		}

	}()
}

func NewRabbitMqClient(nodes []string, userName string, password string, virtualHost string, logger log.Logger, withFunc ...withFunc) *Client {

	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	client := &Client{
		context:       childCtx,
		shutdownFn:    shutdownFn,
		childRoutines: childRoutines,
		consumerBuilder: ConsumerBuilder{
			messageBroker: NewMessageBroker(logger),
		},
		publisherBuilder: &PublisherBuilder{
			messageBroker:    NewMessageBroker(logger),
			startPublisherCh: make(chan bool),
		},
		parameters: MessageBrokerParameter{
			Nodes:           nodes,
			ConcurrentLimit: CONCURRENTLIMIT,
			RetryCount:      RETRYCOUNT,
			PrefetchCount:   PREFECTCOUNT,
			Password:        password,
			UserName:        userName,
			VirtualHost:     virtualHost,
		},
	}

	for _, handler := range withFunc {
		if err := handler(client); err != nil {
			panic(err)
		}
	}
	return client
}

func PrefetchCount(prefetchCount int) withFunc {
	return func(m *Client) error {
		m.parameters.PrefetchCount = prefetchCount
		return nil
	}
}

func RetryCount(retryCount int) withFunc {
	return func(m *Client) error {
		m.parameters.RetryCount = retryCount
		return nil
	}
}
func RetryCountWithWait(retryCount int, retryInterval time.Duration) withFunc {
	return func(m *Client) error {
		m.parameters.RetryCount = retryCount
		m.parameters.RetryInterval = retryInterval
		return nil
	}
}

func (c *Client) Shutdown(reason string) {
	c.shutdownReason = reason
	c.shutdownInProgress = true
	c.shutdownFn()
	c.childRoutines.Wait()
}

func (c *Client) Exit(reason error) int {
	code := 1
	if reason == context.Canceled && c.shutdownReason != "" {
		reason = fmt.Errorf(c.shutdownReason)
		code = 0
	}
	return code
}

func (c *Client) RunConsumers(ctx context.Context) error {

	sendSystemNotification("READY=1")

	c.checkConsumerConnection(ctx)

	for _, consumer := range c.consumerBuilder.consumers {

		consumer := consumer
		c.childRoutines.Go(func() error {

			for {
				select {
				case isConnected := <-consumer.startConsumerCn:

					if isConnected {

						var (
							err           error
							brokerChannel *Channel
						)
						if brokerChannel, err = c.consumerBuilder.messageBroker.CreateChannel(); err != nil {
							panic(err)
						}

						if consumer.consistentExchangeType == ConsistentHashing {

							brokerChannel.createQueue(consumer.queueName).
								exchangeToQueueBind(consumer.consistentExchangeName, consumer.queueName, consumer.consistentRoutingKey, consumer.consistentExchangeType).
								createErrorQueueAndBind(consumer.errorExchangeName, consumer.errorQueueName)

						} else {

							brokerChannel.createQueue(consumer.queueName).
								createErrorQueueAndBind(consumer.errorExchangeName, consumer.errorQueueName)

						}

						for _, item := range consumer.exchanges {

							if consumer.consistentExchangeType == ConsistentHashing {

								brokerChannel.
									createExchange(item.exchangeName, item.exchangeType, item.args).
									exchangeToConsistentExchangeBind(consumer.consistentExchangeName, item.exchangeName, item.routingKey, item.exchangeType)

							} else {
								brokerChannel.
									createExchange(item.exchangeName, item.exchangeType, item.args).
									exchangeToQueueBind(item.exchangeName, consumer.queueName, item.routingKey, item.exchangeType)

							}

						}

						brokerChannel.rabbitChannel.Qos(c.parameters.PrefetchCount, 0, false)

						delivery, _ := brokerChannel.listenToQueue(consumer.queueName)

						if consumer.singleGoroutine {

							c.deliver(brokerChannel, consumer, delivery)

						} else {

							for i := 0; i < c.parameters.PrefetchCount; i++ {
								go func() {
									c.deliver(brokerChannel, consumer, delivery)
								}()
							}
						}
					}
				}
			}

			return nil
		})
	}

	return c.childRoutines.Wait()
}

func (c *Client) AddConsumer(queueName string) *Consumer {

	var consumerDefination = &Consumer{
		queueName:         queueName,
		errorQueueName:    queueName + ERRORPREFIX,
		errorExchangeName: queueName + ERRORPREFIX,
		startConsumerCn:   make(chan bool),
		singleGoroutine:   false,
	}

	var isAlreadyDeclareQueue bool
	for _, item := range c.consumerBuilder.consumers {
		if item.queueName == queueName {
			isAlreadyDeclareQueue = true
		}
	}

	if !isAlreadyDeclareQueue {
		c.consumerBuilder.consumers = append(c.consumerBuilder.consumers, consumerDefination)
	}
	return consumerDefination
}

func (c *Client) AddConsumerWithConsistentHashExchange(queueName string, routingKey string, exchangeName string) *Consumer {

	var consumer = &Consumer{
		queueName:              queueName,
		consistentRoutingKey:   routingKey,
		consistentExchangeType: ConsistentHashing,
		consistentExchangeName: exchangeName,
		errorQueueName:         queueName + ERRORPREFIX,
		errorExchangeName:      queueName + ERRORPREFIX,
		startConsumerCn:        make(chan bool),
		singleGoroutine:        false,
	}

	if routingKey == "" || queueName == "" || exchangeName == "" {
		panic("dont empty value")
	}

	var isAlreadyDeclareQueue bool
	for _, item := range c.consumerBuilder.consumers {
		if item.queueName == queueName {
			isAlreadyDeclareQueue = true
		}
	}

	if !isAlreadyDeclareQueue {
		c.consumerBuilder.consumers = append(c.consumerBuilder.consumers, consumer)
	}
	return consumer
}

func (c *Client) deliver(brokerChannel *Channel, consumer *Consumer, delivery <-chan amqp.Delivery) {
	for d := range delivery {

		Do(func(attempt int) (retry bool, err error) {

			retry = attempt < c.parameters.RetryCount

			defer func() {

				if r := recover(); r != nil {

					if !retry || err == nil {

						err, ok := r.(error)

						if !ok {
							retry = false //Because of panic exception
							err = fmt.Errorf("%v", r)
						}

						stack := make([]byte, 4<<10)
						length := runtime.Stack(stack, false)

						brokerChannel.rabbitChannel.Publish(consumer.errorExchangeName, "", false, false, errorPublishMessage(d.CorrelationId, d.Body, c.parameters.RetryCount, err, fmt.Sprintf("[Exception Recover] %v %s\n", err, stack[:length])))

						select {
						case confirm := <-brokerChannel.notifyConfirm:
							if confirm.Ack {
								break
							}
						case <-time.After(resendDelay):
						}

						d.Ack(false)
					}
					return
				}
			}()
			err = consumer.handleConsumer(Message{CorrelationId: d.CorrelationId, Payload: d.Body, MessageId: d.MessageId, Timestamp: d.Timestamp})
			if err != nil {
				if retry {
					time.Sleep(c.parameters.RetryInterval)
				}
				panic(err)
			}
			d.Ack(false)

			return
		})
	}
}

func (c *Client) AddPublisher(ctx context.Context, exchangeName string, exchangeType ExchangeType, payloads ...interface{}) {

	var payloadTypes []reflect.Type
	for _, item := range payloads {
		if reflect.TypeOf(item).Kind() != reflect.Struct {
			panic(fmt.Sprintf("%s  is not struct", reflect.TypeOf(item).Name()))
		} else {
			payloadTypes = append(payloadTypes, reflect.TypeOf(item))
		}

	}
	var publisher = Publisher{
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		payloadTypes: payloadTypes,
	}

	if len(payloads) == 0 || payloads == nil {
		panic("payloads are not empty")
	}

	for _, item := range c.publisherBuilder.publishers {
		if item.exchangeName == exchangeName {
			panic(fmt.Sprintf("%s exchangename is already declared ", exchangeName))
		}
	}

	c.publisherBuilder.publishers = append(c.publisherBuilder.publishers, publisher)

	c.checkPublisherConnection(ctx)

	c.publisherBuilder.CreateChannel()

}

func (c *Client) Publish(ctx context.Context, routingKey string, payload interface{}) error {

	var (
		publisher *Publisher
	)

	for _, item := range c.publisherBuilder.publishers {

		var tempPublisher = item

		for _, payloadType := range item.payloadTypes {
			if payloadType == reflect.TypeOf(payload) {
				publisher = &tempPublisher
			}
		}
	}

	if publisher == nil {
		return fmt.Errorf("%v is not declared before ", payload)
	}

	return c.publisherBuilder.Publish(ctx, routingKey, publisher.exchangeName, payload)

}

func sendSystemNotification(state string) error {

	notifySocket := os.Getenv("NOTIFY_SOCKET")
	if notifySocket == "" {
		return fmt.Errorf("NOTIFY_SOCKET environment variable empty or unset.")
	}
	socketAddr := &net.UnixAddr{
		Name: notifySocket,
		Net:  "unixgram",
	}
	conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(state))
	conn.Close()
	return err

}
