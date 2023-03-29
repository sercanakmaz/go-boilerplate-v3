package rmqc

import "github.com/streadway/amqp"

type Queue struct {
	Name string
	Arg  map[string]interface{}
}

func newQueue(name string, arg map[string]interface{}) *Queue {
	return &Queue{Name: name, Arg: arg}
}

func (q *Queue) declare(ch *amqp.Channel) error {
	var err error
	_, err = ch.QueueDeclare(q.Name,
		true,
		false,
		false,
		false,
		q.Arg)
	return err
}

func (q *Queue) exchangeBind(ch *amqp.Channel, subscriber *Subscriber) error {
	return ch.QueueBind(q.Name,
		subscriber.routingKey,
		subscriber.exchangeName,
		false,
		nil)
}
