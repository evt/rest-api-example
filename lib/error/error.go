package error

import (
	"fmt"
	"net/http"

	"github.com/evt/rest-api-example/lib/types"
	"github.com/labstack/echo/v4"
)

// HTTPError is our HTTP error definition.
type HTTPError struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"`
}

// Error implements a custom echo error handler that will encode errors as JSON objects rather than
// just return a text body. It will also make sure to not have the redundant information given in
// the echo string encoding of HTTP errors.
func Error(err error, ctx echo.Context) {
	errObj := HTTPError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
	switch err {
	case types.ErrBadRequest:
		errObj.Code = http.StatusBadRequest
	case types.ErrNotFound:
		errObj.Code = http.StatusNotFound
	case types.ErrDuplicateEntry, types.ErrConflict:
		errObj.Code = http.StatusConflict
	case types.ErrForbidden:
		errObj.Code = http.StatusForbidden
	case types.ErrUnprocessableEntity:
		errObj.Code = http.StatusUnprocessableEntity
	case types.ErrPartialOk:
		errObj.Code = http.StatusPartialContent
	case types.ErrGone:
		errObj.Code = http.StatusGone
	case types.ErrUnauthorized:
		errObj.Code = http.StatusUnauthorized
	}
	he, ok := err.(*echo.HTTPError)
	if ok {
		errObj.Code = he.Code
		errObj.Message = fmt.Sprintf("%v", he.Message)
	}
	errObj.Name = http.StatusText(errObj.Code)
	if !ctx.Response().Committed {
		if ctx.Request().Method == echo.HEAD {
			ctx.NoContent(errObj.Code)
		} else {
			ctx.JSON(errObj.Code, errObj)
		}
	}
}
