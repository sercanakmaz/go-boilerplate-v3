package products

import "context"

type (
	IProductService interface {
		AddNew(ctx context.Context, sku string, name string, initialStock int, categoryID int) (*Product, error)
		IncreaseStock(ctx context.Context, sku string, amount int) error
		DecreaseStock(ctx context.Context, sku string, amount int) error
		GetBySku(ctx context.Context, sku string) (*Product, error)
	}
	productService struct {
		Repository IProductRepository
	}
)

func NewProductService(repository IProductRepository) IProductService {
	return &productService{Repository: repository}
}

func (service productService) AddNew(ctx context.Context, sku string, name string, initialStock int, categoryID int) (*Product, error) {
	var product = NewProduct(sku, name, initialStock, categoryID)

	var err = service.Repository.Add(ctx, product)

	return product, err
}
func (service productService) IncreaseStock(ctx context.Context, sku string, amount int) error {
	var (
		err     error
		product *Product
	)

	if product, err = service.Repository.FindOneBySku(ctx, sku); err != nil {
		return err
	}

	product.IncreaseStock(amount)

	return service.Repository.Update(ctx, product)
}
func (service productService) DecreaseStock(ctx context.Context, sku string, amount int) error {
	var (
		err     error
		product *Product
	)

	if product, err = service.Repository.FindOneBySku(ctx, sku); err != nil {
		return err
	}

	product.DecreaseStock(amount)

	return service.Repository.Update(ctx, product)
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
