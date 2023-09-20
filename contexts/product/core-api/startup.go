package core_api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sercanakmaz/go-boilerplate-v3/configs"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/aggregates/products"
	controllers_v1 "github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/controllers/v1"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api/docs"
	product_events "github.com/sercanakmaz/go-boilerplate-v3/events/product/products"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/config"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/rabbitmqv1"
	"github.com/spf13/cobra"
	"net/http"
)

func Init(cmd *cobra.Command, args []string) error {
	docs.Initialize()

	ctx := context.Background()

	var (
		cfg configs.Config
		err error
	)

	if err = config.Load(&cfg); err != nil {
		return err
	}

	var logger = log.NewLogger()
	httpErrorHandler := middlewares.NewHttpErrorHandler()

	e := echo.New()
	e.HTTPErrorHandler = httpErrorHandler.Handler

	e.Use(middlewares.Correlation())
	e.Use(middlewares.Recover())
	e.Use(middlewares.Logger(logger))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/swagger/*", middlewares.WrapHandler)

	messageBus := rabbitmqv1.NewRabbitMqClient(
		[]string{cfg.RabbitMQ.Host},
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		"",
		logger,
		rabbitmqv1.RetryCount(0),
		rabbitmqv1.PrefetchCount(1))

	messageBus.AddPublisher(ctx, "HG.Integration.Product:Created", rabbitmqv1.Topic, product_events.Created{})

	var mongoDb, mongoClient = mongo.NewMongoDb(cfg.Mongo)

	var productRepository = products.NewProductRepository(mongoDb)
	var productService = products.NewProductService(productRepository)

	//e.Use(middleware.BasicAuth(func(username string, password string, ctx echo.Context) (bool, error) {
	//	return userService.AuthUser(context.Background(), username, password)
	//}))

	controllers_v1.NewProductController(e, mongoClient, messageBus, productService, httpErrorHandler)

	return e.Start(fmt.Sprintf(":%v", cfg.Host.Port))
}
