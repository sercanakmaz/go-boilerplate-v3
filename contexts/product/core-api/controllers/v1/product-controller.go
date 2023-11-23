package controllers_v1

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/aggregates/products"
	use_cases "github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/use-cases"
	productModels "github.com/sercanakmaz/go-boilerplate-v3/models/product"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourhttp "github.com/sercanakmaz/go-boilerplate-v3/pkg/http"
	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	string_helper "github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func NewProductController(e *echo.Echo, log logger.Logger, client *mongo.Client, rabbitMQPublisher ddd.IEventPublisher, eventDispatcher ddd.IEventDispatcher, productService products.IProductService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/v1/products")

	CreateProduct(v1, log, client, rabbitMQPublisher, eventDispatcher, productService)
	DeleteProduct(v1, log, client, rabbitMQPublisher, eventDispatcher, productService)

	httpErrorHandler.Add(string_helper.ErrIsNullOrEmpty, http.StatusBadRequest)
	httpErrorHandler.Add(ourhttp.ErrCommandBindFailed, http.StatusBadRequest)
}

// CreateProduct godoc
// @Accept  json
// @Produce  json
// @Param c body product.CreateProductCommand true "body"
// @Success 201 {object} products.Product
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/products/ [post]
func CreateProduct(group *echo.Group, log logger.Logger, client *mongo.Client, rabbitMQPublisher ddd.IEventPublisher, eventDispatcher ddd.IEventDispatcher, productService products.IProductService) {
	group.POST("", func(c echo.Context) error {

		var (
			command *productModels.CreateProductCommand
			product *products.Product
			result  = new(ddd.UseCaseResult[*productModels.CreateProductResponse])
			err     error
		)

		ctx := (c.Get("context")).(context.Context)

		if err = c.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "CreateProductCommand", ourhttp.ErrCommandBindFailed))
		}

		ctx = ddd.NewEventContext(ctx)

		var handler = use_cases.NewCreateProductUseCaseHandler(client, log, productService, rabbitMQPublisher, eventDispatcher)

		if err = ddd.HandleUseCase(ctx, handler, command, result); err != nil {
			panic(fmt.Errorf("%v %w", "CreateProductCommand", ourhttp.ErrUseCaseHandleFailed))
		}

		return c.JSON(http.StatusCreated, product)
	})
}

// DeleteProduct godoc
// @Accept  json
// @Produce  json
// @Param c body product.DeleteProductCommand true "body"
// @Success 200 {object} products.Product
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/products/ [delete]
func DeleteProduct(group *echo.Group, log logger.Logger, client *mongo.Client, rabbitMQPublisher ddd.IEventPublisher, eventDispatcher ddd.IEventDispatcher, productService products.IProductService) {
	group.DELETE("", func(c echo.Context) error {

		var (
			command *productModels.DeleteProductCommand
			result  = new(ddd.UseCaseResult[*productModels.DeleteProductResponse])
			err     error
		)

		ctx := (c.Get("context")).(context.Context)

		if err = c.Bind(&command); err != nil {
			fmt.Printf(err.Error())
			panic(fmt.Errorf("%v %w", "DeleteProductCommand", ourhttp.ErrCommandBindFailed))
		}

		ctx = ddd.NewEventContext(ctx)

		var handler = use_cases.NewDeleteProductUseCaseHandler(client, log, productService, rabbitMQPublisher, eventDispatcher)

		if err = ddd.HandleUseCase(ctx, handler, command, result); err != nil {
			panic(fmt.Errorf("%v %w", "DeleteProductCommand", ourhttp.ErrUseCaseHandleFailed))
		}

		return c.JSON(http.StatusOK, nil)
	})
}
