package order

import (
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orders"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type CreateOrderResponse struct {
	Id          string       `json:"id"`
	OrderNumber string       `json:"orderNumber"`
	Price       shared.Money `json:"price"`
	OrderLines  []OrderLine  `json:"orderLines"`
}

func NewCreateOrderResponse(order *orders.Order, orderLines []*orderlines.OrderLine) *CreateOrderResponse {
	var result = &CreateOrderResponse{
		Id:          order.Id.Hex(),
		OrderNumber: order.OrderNumber,
		Price:       order.Price,
		OrderLines:  NewOrderLines(orderLines),
	}

	return result
}
