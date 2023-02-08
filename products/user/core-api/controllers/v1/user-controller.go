package controllers_v1

import (
	"context"
	"github.com/labstack/echo/v4"
	userModels "go-boilerplate-v3/models/user"
	users2 "go-boilerplate-v3/products/user/aggregates/users"
	"net/http"
)

func CreateGuestUser(group *echo.Group, userService users2.UserService) {
	group.POST("GuestUser", func(ctx echo.Context) error {

		var (
			user *users2.User
			err  error
		)

		if user, err = userService.AddNewGuestUser(context.Background()); err != nil {
			return err
		}

		userModel := userModels.LoadFromUser(user)

		return ctx.JSON(http.StatusCreated, userModel)
	})
}

func GetUserByObjectId(group *echo.Group, userService users2.UserService) {
	group.GET("id/:id", func(ctx echo.Context) error {

		var (
			user *users2.User
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
