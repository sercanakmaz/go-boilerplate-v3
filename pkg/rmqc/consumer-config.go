package rmqc

type ConsumerConfig struct {
	QueueName    string
	ConsumerName string
	Consumer     IConsumer
	Reconnect    Reconnect
	Qos          struct {
		PrefetchCount int
		ConsumerCount int
	}
}
