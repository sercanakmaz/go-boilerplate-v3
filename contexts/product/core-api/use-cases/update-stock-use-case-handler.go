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

type UpdateStockUseCaseHandler struct {
	client            *mongo.Client
	productService    products.IProductService
	rabbitMQPublisher ddd.IEventPublisher
	eventDispatcher   ddd.IEventDispatcher
	logger            logger.Logger
	middlewares       []ddd.IBaseUseCaseMiddleware[*productModels.UpdateStockCommand, *productModels.UpdateStockResponse]
}

func NewUpdateStockUseCaseHandler(client *mongo.Client, log logger.Logger, productService products.IProductService, rabbitMQPublisher ddd.IEventPublisher, eventDispatcher ddd.IEventDispatcher) *UpdateStockUseCaseHandler {
	handler := &UpdateStockUseCaseHandler{
		client:            client,
		productService:    productService,
		rabbitMQPublisher: rabbitMQPublisher,
		eventDispatcher:   eventDispatcher,
		logger:            log,
	}

	handler.SetMiddlewares()

	return handler
}

func (self *UpdateStockUseCaseHandler) SetMiddlewares() {
	self.middlewares = append(self.middlewares, ourMongo.NewTransactionMiddleware[*productModels.UpdateStockCommand, *productModels.UpdateStockResponse](self.client))
	self.middlewares = append(self.middlewares, ddd.NewEventHandlerDispatcherMiddleware[*productModels.UpdateStockCommand, *productModels.UpdateStockResponse](self.eventDispatcher))
	self.middlewares = append(self.middlewares, ddd.NewEventPublisherMiddleware[*productModels.UpdateStockCommand, *productModels.UpdateStockResponse](self.rabbitMQPublisher))
}

func (self *UpdateStockUseCaseHandler) GetMiddlewares() []ddd.IBaseUseCaseMiddleware[*productModels.UpdateStockCommand, *productModels.UpdateStockResponse] {
	return self.middlewares
}

func (self *UpdateStockUseCaseHandler) Handle(ctx context.Context, command *productModels.UpdateStockCommand) (error, *ddd.UseCaseResult[*productModels.UpdateStockResponse]) {

	var (
		err     error
		product *products.Product
	)

	if err = self.productService.UpdateStock(ctx,
		command.Sku,
		command.Stock); err != nil {
		return err, nil
	}

	self.logger.Info(ctx, "product stock updated to the database")

	return nil, ddd.NewUseCaseResultWithContent[*productModels.UpdateStockResponse](product.ConvertUpdateStockResponse())
}
