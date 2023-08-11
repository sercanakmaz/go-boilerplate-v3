package orderlines

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/events/oms/orders"
)

type OrderCancelledEventHandler struct {
	orderLineService IOrderLineService
}

func (self *OrderCancelledEventHandler) Handle(ctx context.Context, created orders.Created) {

}
