package orders

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/event-handler"
)

type Created struct {
	Id          string       `json:"id"`
	OrderNumber string       `json:"orderNumber"`
	Price       shared.Money `json:"price"`
	FinalPrice  shared.Money `json:"finalPrice"`
	event_handler.IBaseEvent
}

func (s *Created) ExchangeName() string {
	return "Orders:Created"
}
