package rmqc

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	config  ConsumerConfig
	channel *amqp.Channel
}

func newConsumer(config ConsumerConfig, channel *amqp.Channel) *Consumer {
	return &Consumer{config: config, channel: channel}
}

func (c *Consumer) Consume(id int) {
	delivery, err := c.channel.Consume(c.config.QueueName,
		fmt.Sprintf("%s (%d/%d)", c.config.ConsumerName, id, c.config.Qos.ConsumerCount),
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Fatal(fmt.Sprintf("CRITICAL: Unable to start consumer %s (%d/%d)", c.config.ConsumerName, id, c.config.Qos.ConsumerCount))
		return
	}

	for d := range delivery {

		context := &ConsumerContext{
			MessageId:     d.MessageId,
			CorrelationId: d.CorrelationId,
			Exchange:      d.Exchange,
			RoutingKey:    d.RoutingKey,
			Data:          d.Body,
		}

		c.config.Consumer.Consume(context)

		d.Ack(false)
	}
}
