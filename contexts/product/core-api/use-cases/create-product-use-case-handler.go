package use_cases

import (
	"context"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/aggregates/products"
	productModels "github.com/sercanakmaz/go-boilerplate-v3/models/product"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	ourMongo "github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateProductUseCaseHandler struct {
	client            *mongo.Client
	productService    products.IProductService
	rabbitMQPublisher ddd.IEventPublisher
	eventDispatcher   ddd.IEventDispatcher
	logger            logger.Logger
	middlewares       []ddd.IBaseUseCaseMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse]
}

func NewCreateProductUseCaseHandler(client *mongo.Client, log logger.Logger, productService products.IProductService, rabbitMQPublisher ddd.IEventPublisher, eventDispatcher ddd.IEventDispatcher) *CreateProductUseCaseHandler {
	handler := &CreateProductUseCaseHandler{
		client:            client,
		productService:    productService,
		rabbitMQPublisher: rabbitMQPublisher,
		eventDispatcher:   eventDispatcher,
		logger:            log,
	}

	handler.SetMiddlewares()

	return handler
}

func (self *CreateProductUseCaseHandler) SetMiddlewares() {
	self.middlewares = append(self.middlewares, ourMongo.NewTransactionMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse](self.client))
	self.middlewares = append(self.middlewares, ddd.NewEventHandlerDispatcherMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse](self.eventDispatcher))
	self.middlewares = append(self.middlewares, ddd.NewEventPublisherMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse](self.rabbitMQPublisher))
}

func (self *CreateProductUseCaseHandler) GetMiddlewares() []ddd.IBaseUseCaseMiddleware[*productModels.CreateProductCommand, *productModels.CreateProductResponse] {
	return self.middlewares
}

func (self *CreateProductUseCaseHandler) Handle(ctx context.Context, command *productModels.CreateProductCommand) (error, *ddd.UseCaseResult[*productModels.CreateProductResponse]) {

	var (
		err     error
		product *products.Product
	)

	if product, err = self.productService.AddNew(ctx,
		command.Sku,
		command.Name,
		command.InitialStock,
		command.Price,
		command.CategoryID); err != nil {
		return err, nil
	}

	self.logger.Info(ctx, "product added to the database")

	return nil, ddd.NewUseCaseResultWithContent[*productModels.CreateProductResponse](product.ConvertCreateProductResponse())
}
