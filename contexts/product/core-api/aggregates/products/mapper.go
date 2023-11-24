package products

import (
	productModels "github.com/sercanakmaz/go-boilerplate-v3/models/product"
)

func (u *Product) ConvertCreateProductResponse() *productModels.CreateProductResponse {
	var result = &productModels.CreateProductResponse{
		Id:         u.Id.Hex(),
		Sku:        u.Sku,
		Name:       u.Name,
		Stock:      u.Stock,
		Price:      u.Price,
		FinalPrice: u.FinalPrice,
		Vat:        u.Vat,
		CategoryId: u.CategoryId,
	}

	return result
}

func (u *Product) ConvertUpdateStockResponse() *productModels.UpdateStockResponse {
	var result = &productModels.UpdateStockResponse{
		Sku:   u.Sku,
		Stock: u.Stock,
	}

	return result
}
