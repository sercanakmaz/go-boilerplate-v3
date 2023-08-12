package orderlines

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/events/oms/orders"
)

type EventHandler struct {
	orderLineService IOrderLineService
}

func (s *EventHandler) Handle(ctx context.Context, paymentRejected orders.PaymentRejected) error {
	var (
		err        error
		orderLines []*OrderLine
	)
	if orderLines, err = s.orderLineService.FindByOrderNumber(ctx, paymentRejected.OrderNumber); err != nil {
		return err
	}

	for _, orderLine := range orderLines {
		if _, err = s.orderLineService.Cancel(ctx, orderLine.Id.Hex(), "OrderPaymentRejected"); err != nil {
			return err
		}
	}

	return nil
}
