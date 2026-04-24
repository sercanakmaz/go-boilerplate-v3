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

type DeleteProductUseCaseHandler struct {
	ddd.BaseUseCaseHandler[*productModels.DeleteProductCommand, *productModels.DeleteProductResponse]
	client            *mongo.Client
	productService    products.IProductService
	rabbitMQPublisher ddd.IEventPublisher
	eventDispatcher   ddd.IEventDispatcher
	logger            logger.Logger
}

func NewDeleteProductUseCaseHandler(client *mongo.Client, log logger.Logger, productService products.IProductService, rabbitMQPublisher ddd.IEventPublisher, eventDispatcher ddd.IEventDispatcher) *DeleteProductUseCaseHandler {
	handler := &DeleteProductUseCaseHandler{
		client:            client,
		productService:    productService,
		rabbitMQPublisher: rabbitMQPublisher,
		eventDispatcher:   eventDispatcher,
		logger:            log,
	}

	handler.SetMiddlewares()

	return handler
}

func (self *DeleteProductUseCaseHandler) SetMiddlewares() {
	self.Use(
		ourMongo.NewTransactionMiddleware[*productModels.DeleteProductCommand, *productModels.DeleteProductResponse](self.client),
		ddd.NewEventHandlerDispatcherMiddleware[*productModels.DeleteProductCommand, *productModels.DeleteProductResponse](self.eventDispatcher),
		ddd.NewEventPublisherMiddleware[*productModels.DeleteProductCommand, *productModels.DeleteProductResponse](self.rabbitMQPublisher),
	)
}

func (self *DeleteProductUseCaseHandler) Handle(ctx context.Context, command *productModels.DeleteProductCommand) (error, *ddd.UseCaseResult[*productModels.DeleteProductResponse]) {

	var (
		err error
	)

	if err = self.productService.Delete(ctx, command.Sku); err != nil {
		return err, nil
	}

	self.logger.Info(ctx, "product deleted from the database")

	return nil, ddd.NewUseCaseResultWithContent[*productModels.DeleteProductResponse](nil)
}
