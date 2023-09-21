package controllers_v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/aggregates/products"
	use_cases "github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/use-cases"
	productModels "github.com/sercanakmaz/go-boilerplate-v3/models/product"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourhttp "github.com/sercanakmaz/go-boilerplate-v3/pkg/http"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	string_helper "github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func NewProductController(e *echo.Echo, client *mongo.Client, rabbitMQPublisher ddd.IEventPublisher, productService products.IProductService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/v1/products")

	CreateProduct(v1, client, rabbitMQPublisher, productService)

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
func CreateProduct(group *echo.Group, client *mongo.Client, rabbitMQPublisher ddd.IEventPublisher, productService products.IProductService) {
	group.POST("", func(eCtx echo.Context) error {

		var (
			command *productModels.CreateProductCommand
			product *products.Product
			result  = new(ddd.UseCaseResult[*productModels.CreateProductResponse])
			err     error
		)

		if err = eCtx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "CreateProductCommand", ourhttp.ErrCommandBindFailed))
		}

		var handler = use_cases.NewCreateProductUseCaseHandler(client, productService)

		ctx := ddd.NewEventContext(eCtx.Request().Context())

		if err = ddd.HandleUseCase(ctx, handler, command, result); err != nil {
			panic(fmt.Errorf("%v %w", "CreateProductCommand", ourhttp.ErrUseCaseHandleFailed))
		}

		if err = rabbitMQPublisher.Publish(ctx); err != nil {
			panic(fmt.Errorf("%v %w", "CreateProductCommand", ourhttp.ErrRabbitMQPublishFailed))
		}

		return eCtx.JSON(http.StatusCreated, product)
	})
}
