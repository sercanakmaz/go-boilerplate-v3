package controllers_v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/user/core-api/aggregates/users"
	userModels "github.com/sercanakmaz/go-boilerplate-v3/models/user"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/middlewares"
	string_helper "github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"net/http"
)

func NewUserController(e *echo.Echo, userService users.UserService, httpErrorHandler middlewares.HttpErrorHandler) {
	v1 := e.Group("/users/v1/")

	CreateGuestUser(v1, userService)
	GetUserByObjectId(v1, userService)

	httpErrorHandler.Add(string_helper.ErrIsNullOrEmpty, http.StatusBadRequest)
	httpErrorHandler.Add(users.ErrAlreadyExistRole, http.StatusConflict)
}

func CreateGuestUser(group *echo.Group, userService users.UserService) {
	group.POST("GuestUser", func(ctx echo.Context) error {

		var (
			user *users.User
			err  error
		)

		if user, err = userService.AddNewGuestUser(context.Background()); err != nil {
			return err
		}

		userModel := userModels.LoadFromUser(user)

		return ctx.JSON(http.StatusCreated, userModel)
	})
}

func GetUserByObjectId(group *echo.Group, userService users.UserService) {
	group.GET("id/:id", func(ctx echo.Context) error {

		var (
			user *users.User
			err  error
		)

		id := ctx.Param("id")
		if user, err = userService.GetUserById(context.Background(), id); err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		userModel := userModels.LoadFromUser(user)

		return ctx.JSON(http.StatusCreated, userModel)
	})
}
