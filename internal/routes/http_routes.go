package routes

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(w *HTTPWrapper) func(e *echo.Group) {
	return func(e *echo.Group) {
		initViewRoutes(e, w)
		api := e.Group("api")
		initAPIRoutes(api, w)
	}
}

func initViewRoutes(e *echo.Group, w *HTTPWrapper) {
	e.GET("/vehicles", w.GetVehiclesPage)
	e.GET("/form/create-fuel", w.GetFormCreateFuel)
	e.GET("/form/create-odometer", w.GetFormCreateOdometer)
}

func initAPIRoutes(e *echo.Group, w *HTTPWrapper) {
	e.POST("/vehicles", w.CreateVehicle)
	e.GET("/vehicles/:vehicle_id", w.GetVehicle)
}
