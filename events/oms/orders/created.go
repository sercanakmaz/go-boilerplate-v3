package orders

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type Created struct {
	Id          string       `json:"id"`
	OrderNumber string       `json:"orderNumber"`
	Price       shared.Money `json:"price"`
	FinalPrice  shared.Money `json:"finalPrice"`
}

func (s *Created) ExchangeName() string {
	return "Orders:Created"
}
