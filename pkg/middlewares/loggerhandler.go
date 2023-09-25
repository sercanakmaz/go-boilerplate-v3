package middlewares

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"time"

	logger "github.com/sercanakmaz/go-boilerplate-v3/pkg/log"
)

const LogExtra = "Log_Extra"

type Stats struct {
	Statuses map[string]int `json:"statuses"`
	Uptime   time.Time      `json:"uptime"`
}

func Logger(log logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var (
				start    time.Time
				contents []byte
				field    logger.Field
			)
			if c.Request().Body != nil {
				contents, _ = ioutil.ReadAll(c.Request().Body)
				c.Request().Body = ioutil.NopCloser(bytes.NewReader(contents))
			}

			requestBodySizeInKb := binary.Size(contents) / 1024
			if requestBodySizeInKb > 250 {
				contents = []byte("Request body is too big. Ignored")
			}

			var logExtraRaw = c.Get(LogExtra)
			var logExtra map[string]interface{}

			if logExtraRaw != nil {
				logExtra = logExtraRaw.(map[string]interface{})
			}

			start = time.Now()
			field = logger.Field{
				Message:     "Starting controller request",
				RequestBody: string(contents),
				HostName:    c.Request().Host,
				Url:         string(c.Request().RequestURI),
				HttpMethod:  c.Request().Method,
				Extra:       logExtra,
			}

			var correlationId = (c.Get("correlationId")).(string)

			// Bind correlation id to new context
			ctx := context.WithValue(c.Request().Context(), "correlationId", correlationId)

			// Bind logger with correlation id
			ctx = log.WithCorrelationId(ctx, correlationId)

			c.Set("context", ctx)

			log.Request(ctx, field)

			if err := next(c); err != nil {

				var echoError, ok = err.(*echo.HTTPError)
				if ok {
					field = logger.Field{
						Message:        echoError.Error(),
						HttpStatusCode: echoError.Code,
						ResponseBody:   echoError.Error(),
						Duration:       time.Since(start).Milliseconds(),
						Url:            string(c.Request().RequestURI),
						HostName:       c.Request().Host,
						HttpMethod:     c.Request().Method,
						Extra:          logExtra,
					}
					log.Response(ctx, field)
				}

				return err
			} else {

				var logExtraRaw = c.Get(LogExtra)
				var logExtra map[string]interface{}

				if logExtraRaw != nil {
					logExtra = logExtraRaw.(map[string]interface{})
				}

				field = logger.Field{
					Message:        "Finished controller request",
					HttpStatusCode: c.Response().Status,
					ResponseBody:   "ignored",
					Duration:       time.Since(start).Milliseconds(),
					Url:            string(c.Request().RequestURI),
					HostName:       c.Request().Host,
					HttpMethod:     c.Request().Method,
					Extra:          logExtra,
				}

				log.Response(ctx, field)
				return nil
			}

		}
	}
}

func BindEchoContext(c echo.Context, logger logger.Logger) context.Context {
	var correlationId = (c.Get("correlationId")).(string)
	var ctx = context.WithValue(context.Background(), "correlationId", correlationId)

	ctx = context.WithValue(ctx, "echoContext", c)

	return logger.WithCorrelationId(ctx, correlationId)
}
