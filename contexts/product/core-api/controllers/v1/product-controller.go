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
	GetProductBySku(v1, productService)
	IncreaseStock(v1, productService)
	DecreaseStock(v1, productService)

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

		if product, err = productService.AddNew(ctx.Request().Context(),
			command.Sku,
			command.Name,
			command.InitialStock,
			command.Price,
			command.CategoryID); err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, product)
	})
}

// IncreaseStock godoc
// @Accept  json
// @Produce  json
// @Param sku path string true "SKU"
// @Param c body product.IncreaseProductStockCommand true "body"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/products/increase-stock/{sku} [put]
func IncreaseStock(group *echo.Group, productService products.IProductService) {
	group.PUT("increase-stock/:sku", func(ctx echo.Context) error {

		var (
			command *productModels.IncreaseProductStockCommand
			err     error
		)

		if err = ctx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "IncreaseProductStockCommand", ourhttp.ErrCommandBindFailed))
		}

		var sku = ctx.Param("sku")

		if err = productService.IncreaseStock(ctx.Request().Context(),
			sku,
			command.Stock); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	})
}

// DecreaseStock godoc
// @Accept  json
// @Produce  json
// @Param sku path string true "SKU"
// @Param c body product.DecreaseProductStockCommand true "body"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/products/decrease-stock/{sku} [put]
func DecreaseStock(group *echo.Group, productService products.IProductService) {
	group.PUT("decrease-stock/:sku", func(ctx echo.Context) error {

		var (
			command *productModels.DecreaseProductStockCommand
			err     error
		)

		if err = ctx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "DecreaseProductStockCommand", ourhttp.ErrCommandBindFailed))
		}

		var sku = ctx.Param("sku")

		if err = productService.DecreaseStock(ctx.Request().Context(),
			sku,
			command.Stock); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	})
}

// GetProductBySku godoc
// @Accept  json
// @Produce  json
// @Param sku path string true "SKU"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/products/id/{id} [get]
func GetProductBySku(group *echo.Group, productService products.IProductService) {
	group.GET("id/:id", func(ctx echo.Context) error {

		var (
			product *products.Product
			err     error
		)

		id := ctx.Param("id")
		if product, err = productService.GetBySku(ctx.Request().Context(), id); err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusCreated, product)
	})
}
