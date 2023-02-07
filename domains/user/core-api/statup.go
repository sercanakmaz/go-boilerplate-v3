package core_api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-boilerplate-v3/configs"
	users2 "go-boilerplate-v3/domains/user/aggregates/users"
	usersController "go-boilerplate-v3/domains/user/core-api/controllers/v1"
	"go-boilerplate-v3/pkg/config"
	"go-boilerplate-v3/pkg/log"
	"go-boilerplate-v3/pkg/middlewares"
	string_helper "go-boilerplate-v3/pkg/string-helper"
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

	NewUserController(e, userService, httpErrorHandler)

	if err := e.Start(fmt.Sprintf(":%v", cfg.Host.Port)); err != nil {
		panic(err)
	}
}

func NewUserController(e *echo.Echo, userService users2.UserService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/users/v1/")

	usersController.CreateGuestUser(v1, userService)
	usersController.GetUserByObjectId(v1, userService)

	httpErrorHandler.Add(string_helper.ErrIsNullOrEmpty, http.StatusBadRequest)
	httpErrorHandler.Add(users2.ErrAlreadyExistRole, http.StatusConflict)
}
