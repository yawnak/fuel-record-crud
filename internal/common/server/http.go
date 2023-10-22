package server

import "github.com/labstack/echo/v4"

type HTTP struct {
	e *echo.Echo
}

func NewHTTP(e *echo.Echo) (*HTTP, error) {
	return &HTTP{e: e}, nil
}

func (h *HTTP) ListenAndServe(port string) error {
	return h.e.Start(":" + port)
}
