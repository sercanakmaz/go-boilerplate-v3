package controllers_v1

import (
	"github.com/labstack/echo/v4"
	"go-boilerplate-v3/domains/product/core-api/aggregates/products"
	users2 "go-boilerplate-v3/domains/user/aggregates/users"
	productModels "go-boilerplate-v3/models/product"
	"go-boilerplate-v3/pkg/middlewares"
	string_helper "go-boilerplate-v3/pkg/string-helper"
	"net/http"
)

func NewProductController(e *echo.Echo, productService products.IProductService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/v1/products/")

	CreateProduct(v1, productService)
	GetProductBySku(v1, productService)

	httpErrorHandler.Add(string_helper.ErrIsNullOrEmpty, http.StatusBadRequest)
	httpErrorHandler.Add(users2.ErrAlreadyExistRole, http.StatusConflict)
}

func CreateProduct(group *echo.Group, productService products.IProductService) {
	group.POST("", func(ctx echo.Context) error {

		var (
			createProductCommand *productModels.CreateProductCommand
			product              *products.Product
			err                  error
		)

		if err := ctx.Bind(&createProductCommand); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		if product, err = productService.AddNew(ctx.Request().Context(),
			createProductCommand.Sku,
			createProductCommand.Name,
			createProductCommand.InitialStock,
			createProductCommand.CategoryID); err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, product)
	})
}

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
