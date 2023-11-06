package server

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yawnak/fuel-record-crud/web"
)

type HTTP struct {
	e *echo.Echo
}

func RunHTTPServer(addr string, renderer echo.Renderer, createHandler func(router *echo.Group)) error {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = renderer
	fs := echo.MustSubFS(web.Dist, "dist")
	e.StaticFS("/dist", fs)
	rootRouter := e.Group("")
	createHandler(rootRouter)

	saveRoutes(e)
	return e.Start(":" + addr)
}

func (h *HTTP) ListenAndServe(port string) error {
	return h.e.Start(":" + port)
}

func saveRoutes(e *echo.Echo) error {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("api/echo-routes.json", data, 0644)
}
