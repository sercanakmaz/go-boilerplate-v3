package products

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
)

type (
	IProductService interface {
		AddNew(ctx context.Context, sku string, name string, initialStock int, price shared.Money, categoryID int) (*Product, error)
		UpdateStock(ctx context.Context, sku string, stock int) error
		Delete(ctx context.Context, sku string) error
		GetBySku(ctx context.Context, sku string) (*Product, error)
	}
	productService struct {
		Repository IProductRepository
	}
)

func NewProductService(repository IProductRepository) IProductService {
	return &productService{Repository: repository}
}

func (service productService) AddNew(ctx context.Context, sku string, name string, initialStock int, price shared.Money, categoryID int) (*Product, error) {
	var product = NewProduct(sku, name, initialStock, price, categoryID)

	product.brandCross()

	var err = service.Repository.Add(ctx, product)

	return product, err
}

func (service productService) UpdateStock(ctx context.Context, sku string, stock int) error {
	var product = UpdateProductStock(sku, stock)

	var err = service.Repository.UpdateStock(ctx, product)

	return err
}

func (service productService) Delete(ctx context.Context, sku string) error {
	var (
		err     error
		product *Product
	)

	if product, err = service.Repository.FindOneBySku(ctx, sku); err != nil {
		return err
	}

	product.deleteProduct()

	return service.Repository.Delete(ctx, product)
}

func (service productService) GetBySku(ctx context.Context, sku string) (*Product, error) {
	var (
		err     error
		product *Product
	)

	if product, err = service.Repository.FindOneBySku(ctx, sku); err != nil {
		return nil, err
	}

	return product, err
}
