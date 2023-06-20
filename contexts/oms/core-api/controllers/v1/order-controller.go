package controllers_v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orders"
	orderModels "github.com/sercanakmaz/go-boilerplate-v3/models/order"
	ourhttp "github.com/sercanakmaz/go-boilerplate-v3/pkg/http"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	string_helper "github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"net/http"
)

func NewOrderController(e *echo.Echo, orderService orders.IOrderService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/v1/orders/")

	CreateOrder(v1, orderService)
	GetOrderByOrderNumber(v1, orderService)

	httpErrorHandler.Add(string_helper.ErrIsNullOrEmpty, http.StatusBadRequest)
	httpErrorHandler.Add(ourhttp.ErrCommandBindFailed, http.StatusBadRequest)
}

// CreateOrder godoc
// @Accept  json
// @Produce  json
// @Param c body order.CreateOrderCommand true "body"
// @Success 201 {object} orders.Order
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/orders/ [post]
func CreateOrder(group *echo.Group, orderService orders.IOrderService) {
	group.POST("", func(ctx echo.Context) error {

		var (
			command *orderModels.CreateOrderCommand
			order   *orders.Order
			err     error
		)

		if err = ctx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "CreateOrderCommand", ourhttp.ErrCommandBindFailed))
		}

		if order, err = orderService.AddNew(ctx.Request().Context(),
			command.OrderNumber,
			command.Price); err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, order)
	})
}

// GetOrderByOrderNumber godoc
// @Accept  json
// @Produce  json
// @Param orderNumber path string true "Order Number"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/orders/orderNumber/{orderNumber} [get]
func GetOrderByOrderNumber(group *echo.Group, orderService orders.IOrderService) {
	group.GET("orderNumber/:orderNumber", func(ctx echo.Context) error {

		var (
			order *orders.Order
			err   error
		)

		id := ctx.Param("id")
		if order, err = orderService.GetByOrderNumber(ctx.Request().Context(), id); err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusCreated, order)
	})
}
