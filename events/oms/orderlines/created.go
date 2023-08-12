package orderlines

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type Created struct {
	Id          string       `json:"id"`
	Sku         string       `json:"sku"`
	OrderNumber string       `json:"orderNumber"`
	Price       shared.Money `json:"price"`
	ddd.IBaseEvent
}

func (s *Created) ExchangeName() string {
	return "OrderLines:Created"
}
