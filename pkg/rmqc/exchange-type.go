package rmqc

type ExchangeType string

const (
	ExchangeType_Topic   ExchangeType = "topic"
	ExchangeType_Direct  ExchangeType = "direct"
	ExchangeType_Fanout  ExchangeType = "fanout"
	ExchangeType_Headers ExchangeType = "headers"
)
