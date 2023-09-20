package use_cases

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/aggregates/products"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/infra"
	productModels "github.com/sercanakmaz/go-boilerplate-v3/models/product"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourMongo "github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateProductUseCaseHandler struct {
	client         *mongo.Client
	productService products.IProductService
	middlewares    []ddd.IBaseUseCaseMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse]
}

func NewCreateProductUseCaseHandler(client *mongo.Client, productService products.IProductService) *CreateProductUseCaseHandler {
	var handler = &CreateProductUseCaseHandler{
		client:         client,
		productService: productService,
	}

	handler.middlewares = append(handler.middlewares, ourMongo.NewTransactionMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse](client))
	handler.middlewares = append(handler.middlewares, ddd.NewEventHandlerDispatcherMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse](infra.NewEventHandlerDispatcher()))

	return handler
}

func (self *CreateProductUseCaseHandler) GetMiddlewares() []ddd.IBaseUseCaseMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse] {
	return self.middlewares
}

func (self *CreateProductUseCaseHandler) Handle(ctx context.Context, command *productModels.CreateProductCommand) (error, *ddd.UseCaseResult[*productModels.CreateProductResponse]) {

	var (
		err     error
		product *products.Product
	)

	// TODO: Sercan'a sor! Command mı dto mu?
	if product, err = self.productService.AddNew(ctx, command); err != nil {
		return err, nil
	}

	return nil, ddd.NewUseCaseResultWithContent[*productModels.CreateProductResponse](productModels.NewCreateProductResponse(product))
}
