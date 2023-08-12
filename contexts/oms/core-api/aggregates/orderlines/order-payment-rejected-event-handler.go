package orderlines

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/events/oms/orders"
)

type EventHandler struct {
	orderLineService IOrderLineService
}

func (s *EventHandler) Handle(ctx context.Context, paymentRejected orders.PaymentRejected) {

}
