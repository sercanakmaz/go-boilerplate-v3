package orderlines

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type Cancelled struct {
	Id           string `json:"id"`
	Sku          string `json:"sku"`
	OrderNumber  string `json:"orderNumber"`
	Status       string `json:"status"`
	CancelReason string `json:"cancelReason"`
	ddd.IBaseEvent
}

func (s *Cancelled) ExchangeName() string {
	return "OrderLines:Cancelled"
}
