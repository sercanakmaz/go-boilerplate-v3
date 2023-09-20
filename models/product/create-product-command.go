package product

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type CreateProductCommand struct {
	Sku          string       `json:"sku"`
	Name         string       `json:"name"`
	Brand        string       `json:"brand"`
	InitialStock int          `json:"initialStock"`
	CategoryID   int          `json:"categoryID"`
	Price        shared.Money `json:"price"`
}
