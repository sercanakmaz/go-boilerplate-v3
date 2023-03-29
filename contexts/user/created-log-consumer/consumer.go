package created_log_consumer

import (
	"fmt"
	"github.com/rozturac/rmqc"
	"go-boilerplate-v3/events/user/users"
)

type UserCreatedConsumer struct {
}

func NewUserCreatedConsumer() *UserCreatedConsumer {
	return &UserCreatedConsumer{}
}

func (u UserCreatedConsumer) Configure(builder *rmqc.ConsumerBuilder) {
	builder.BindQueue("user-created-logs")
	builder.SubscribeAsTopic("UserCreated", "")
	builder.SetConsumerName("user-created-logs-consumer")
	builder.SetPrefetchCount(3)
	builder.SetConsumerCount(3)
}

func (u UserCreatedConsumer) Consume(context *rmqc.ConsumerContext) {
	var event users.UserCreated
	context.Unmarshal(&event)
	fmt.Println(event)
}
