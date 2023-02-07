package middlewares

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func Correlation() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var guid = uuid.NewV4()
			c.Set("correlationId", guid.String())
			return next(c)
		}
	}
}
