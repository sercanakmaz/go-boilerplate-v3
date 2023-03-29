package rmqc

type IConsumer interface {
	Configure(builder *ConsumerBuilder)
	Consume(context *ConsumerContext)
}
