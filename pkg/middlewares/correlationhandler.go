package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Correlation() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var guid = uuid.New()
			c.Set("correlationId", guid.String())
			return next(c)
		}
	}
}
