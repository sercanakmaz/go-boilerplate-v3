package order

import (
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type OrderLine struct {
	Id    string       `json:"id"`
	Sku   string       `json:"sku"`
	Price shared.Money `json:"price"`
}

func NewOrderLine(line *orderlines.OrderLine) OrderLine {
	return OrderLine{
		Id:    line.Id.Hex(),
		Sku:   line.Sku,
		Price: line.Price,
	}
}

func NewOrderLines(orderLines []*orderlines.OrderLine) []OrderLine {
	var result []OrderLine

	for _, line := range orderLines {
		result = append(result, NewOrderLine(line))
	}

	return result
}
