package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
)

type ErrAdapter interface {
	Adapt(err error) *httperr.SError
}

func NewErrorInterceptorMiddleware(adapter ErrAdapter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return adapter.Adapt(next(c))
		}
	}
}
