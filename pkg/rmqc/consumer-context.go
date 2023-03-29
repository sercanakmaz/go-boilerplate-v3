package rmqc

import (
	"encoding/json"
)

type ConsumerContext struct {
	MessageId     string
	CorrelationId string
	Exchange      string
	RoutingKey    string
	Data          []byte
}

func (c *ConsumerContext) Unmarshal(t interface{}) {
	json.Unmarshal(c.Data, t)
}
