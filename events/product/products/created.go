package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/event-handler"
)

type Created struct {
	Id         string       `json:"id"`
	Sku        string       `json:"sku"`
	Name       string       `json:"name"`
	Stock      int          `json:"stock"`
	Price      shared.Money `json:"price"`
	FinalPrice shared.Money `json:"finalPrice"`
	CategoryId int          `json:"categoryId"`
	event_handler.IBaseEvent
}

func (s *Created) ExchangeName() string {
	return "Products:Created"
}
