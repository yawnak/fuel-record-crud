package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yawnak/fuel-record-crud/internal/adapters/erradapt"
)

func ErrAdapter(next echo.HandlerFunc) echo.HandlerFunc {
	adapter := erradapt.NewAdapter()
	return func(c echo.Context) error {
		return adapter.Adapt(next(c))
	}
}
