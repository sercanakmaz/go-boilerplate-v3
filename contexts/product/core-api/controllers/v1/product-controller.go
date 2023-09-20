package controllers_v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/aggregates/products"
	productModels "github.com/sercanakmaz/go-boilerplate-v3/models/product"
	ourhttp "github.com/sercanakmaz/go-boilerplate-v3/pkg/http"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	string_helper "github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"net/http"
)

func NewProductController(e *echo.Echo, productService products.IProductService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/v1/products/")

	CreateProduct(v1, productService)

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
func CreateProduct(group *echo.Group, productService products.IProductService) {
	group.POST("", func(ctx echo.Context) error {

		var (
			command *productModels.CreateProductCommand
			product *products.Product
			err     error
		)

		if err = ctx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "CreateProductCommand", ourhttp.ErrCommandBindFailed))
		}

		// TODO: Sercan'a sor! Command mı dto mu?
		if product, err = productService.AddNew(ctx.Request().Context(), command); err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, product)
	})
}
