package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"runtime"
)

type (
	// RecoverConfig defines the config for Recover middleware.
	RecoverConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Size of the stack to be printed.
		// Optional. Default value 4KB.
		StackSize int `yaml:"stack_size"`

		// DisableStackAll disables formatting stack traces of all other goroutines
		// into buffer after the trace for the current goroutine.
		// Optional. Default value false.
		DisableStackAll bool `yaml:"disable_stack_all"`

		// DisablePrintStack disables printing stack trace.
		// Optional. Default value as false.
		DisablePrintStack bool `yaml:"disable_print_stack"`

		// LogLevel is log level to printing stack trace.
		// Optional. Default value 0 (Print).
		LogLevel log.Lvl
	}
)

var (
	// DefaultRecoverConfig is the default Recover middleware config.
	DefaultRecoverConfig = RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
	}
)

// Recover returns a middleware which recovers from panics anywhere in the chain
// and handles the control to the centralized HTTPErrorHandler.
func Recover() echo.MiddlewareFunc {
	return RecoverWithConfig(DefaultRecoverConfig)
}

// RecoverWithConfig returns a Recover middleware with config.
// See: `Recover()`.
func RecoverWithConfig(config RecoverConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultRecoverConfig.Skipper
	}
	if config.StackSize == 0 {
		config.StackSize = DefaultRecoverConfig.StackSize
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)
					if !config.DisablePrintStack {
						msg := log.JSON{
							"message": fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length]),
							"id":      c.Response().Header().Get(echo.HeaderXRequestID),
						}
						switch config.LogLevel {
						case log.DEBUG:
							c.Logger().Debugj(msg)
						case log.INFO:
							c.Logger().Infoj(msg)
						case log.WARN:
							c.Logger().Warnj(msg)
						case log.ERROR:
							c.Logger().Errorj(msg)
						case log.OFF:
							// None.
						default:
							c.Logger().Printj(msg)
						}
					}
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
