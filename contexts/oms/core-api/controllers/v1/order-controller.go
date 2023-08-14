package controllers_v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/aggregates/orders"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api/use-cases"
	orderModels "github.com/sercanakmaz/go-boilerplate-v3/models/order"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	ourhttp "github.com/sercanakmaz/go-boilerplate-v3/pkg/http"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	string_helper "github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type OrderController struct {
}

func NewOrderController(e *echo.Echo, client *mongo.Client, orderService orders.IOrderService, orderLineService orderlines.IOrderLineService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/v1/orders/")

	CreateOrder(v1, client, orderService, orderLineService)
	GetOrderByOrderNumber(v1, client, orderService, orderLineService)
	RejectPayment(v1, client, orderService, orderLineService)

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
func CreateOrder(group *echo.Group, client *mongo.Client, orderService orders.IOrderService, orderLineService orderlines.IOrderLineService) {
	group.POST("", func(ctx echo.Context) error {

		var (
			command *orderModels.CreateOrderCommand
			result  = new(ddd.UseCaseResult[*orderModels.CreateOrderResponse])
			err     error
		)

		if err = ctx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "CreateOrderCommand", ourhttp.ErrCommandBindFailed))
		}

		var handler = use_cases.NewCreateOrderUseCaseHandler(client, orderService, orderLineService)

		if err = ddd.HandleUseCase(ctx.Request().Context(), handler, command, result); err != nil {
			panic(fmt.Errorf("%v %w", "CreateOrderCommand", ourhttp.ErrUseCaseHandleFailed))
		}

		return ctx.JSON(http.StatusCreated, result.Content)
	})
}

// RejectPayment godoc
// @Accept  json
// @Produce  json
// @Param orderNumber path string true "Order Number"
// @Param c body order.RejectOrderPaymentCommand true "body"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v1/orders/{orderNumber}/reject-payment [put]
func RejectPayment(group *echo.Group, client *mongo.Client, orderService orders.IOrderService, orderLineService orderlines.IOrderLineService) {
	group.PUT(":orderNumber/reject-payment", func(ctx echo.Context) error {

		var (
			command *orderModels.RejectOrderPaymentCommand
			result  = new(ddd.UseCaseResult[string])
			err     error
		)

		if err = ctx.Bind(&command); err != nil {
			panic(fmt.Errorf("%v %w", "RejectOrderPaymentCommand", ourhttp.ErrCommandBindFailed))
		}

		command.OrderNumber = ctx.Param("orderNumber")

		var handler = use_cases.NewRejectOrderPaymentUseCaseHandler(client, orderService, orderLineService)

		if err = ddd.HandleUseCase(ctx.Request().Context(), handler, command, result); err != nil {
			panic(fmt.Errorf("%v %w", "RejectOrderPaymentCommand", ourhttp.ErrUseCaseHandleFailed))
		}

		return ctx.String(http.StatusNoContent, result.Content)
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
func GetOrderByOrderNumber(group *echo.Group, client *mongo.Client, orderService orders.IOrderService, orderLineService orderlines.IOrderLineService) {
	group.GET("orderNumber/:orderNumber", func(ctx echo.Context) error {

		var (
			order *orders.Order
			err   error
		)

		orderNumber := ctx.Param("orderNumber")
		if order, err = orderService.GetByOrderNumber(ctx.Request().Context(), orderNumber); err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusCreated, order)
	})
}
