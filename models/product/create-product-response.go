package product

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type CreateProductResponse struct {
	Id         string       `json:"id"`
	Sku        string       `json:"sku"`
	Name       string       `json:"name"`
	Stock      int          `json:"stock"`
	Price      shared.Money `json:"price"`
	FinalPrice shared.Money `json:"finalPrice"`
	Vat        float64      `json:"vat"`
	CategoryId int          `json:"categoryId"`
}
