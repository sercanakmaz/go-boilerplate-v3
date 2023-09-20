package ddd

type IBaseEvent interface {
	EventName() string
	ExchangeName() string
}
