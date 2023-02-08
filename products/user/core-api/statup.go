package core_api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-boilerplate-v3/configs"
	"go-boilerplate-v3/pkg/config"
	"go-boilerplate-v3/pkg/log"
	"go-boilerplate-v3/pkg/middlewares"
	users2 "go-boilerplate-v3/products/user/aggregates/users"
	usersController "go-boilerplate-v3/products/user/core-api/controllers/v1"
	"net/http"
)

func Init() {

	var (
		cfg configs.Config
		err error
	)

	if err = config.Load(&cfg); err != nil {
		panic(err)
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

	var userService = users2.NewUserServiceResolve(cfg.RabbitMQ, cfg.Mongo)

	//e.Use(middleware.BasicAuth(func(username string, password string, ctx echo.Context) (bool, error) {
	//	return userService.AuthUser(context.Background(), username, password)
	//}))

	usersController.NewUserController(e, userService, httpErrorHandler)

	if err := e.Start(fmt.Sprintf(":%v", cfg.Host.Port)); err != nil {
		panic(err)
	}
}
