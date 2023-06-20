package order

import "github.com/sercanakmaz/go-boilerplate-v3/models/shared"

type CreateOrderCommand struct {
	OrderNumber string       `json:"orderNumber"`
	Price       shared.Money `json:"price"`
}
