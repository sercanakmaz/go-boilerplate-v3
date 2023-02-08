package middlewares

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	HttpErrorHandler interface {
		Add(err error, code int)
		getStatusCode(err error) int
	}
	httpErrorHandler struct {
		statusCodes map[error]int
	}
)

func NewHttpErrorHandler() *httpErrorHandler {
	var result = &httpErrorHandler{}
	result.statusCodes = make(map[error]int)
	return result
}

func (self *httpErrorHandler) Add(err error, code int) {
	self.statusCodes[err] = code
}

func (self *httpErrorHandler) getStatusCode(err error) int {
	for key, value := range self.statusCodes {
		if errors.Is(err, key) {
			return value
		}
	}

	return http.StatusInternalServerError
}

func unwrapRecursive(err error) error {
	var originalErr = err

	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}

		originalErr = internalErr
	}

	return originalErr
}

func (self *httpErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    self.getStatusCode(err),
			Message: unwrapRecursive(err).Error(),
		}
	}

	// {
	//   Code: "ErrIsNullOrEmpty"
	//   Message: "role ErrIsNullOrEmpty"
	// }

	code := he.Code
	message := he.Message
	if m, ok := he.Message.(string); ok {
		message = map[string]interface{}{"code": m, "message": err.Error()}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
