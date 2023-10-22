package routes

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, w *HTTPWrapper) *echo.Echo {
	if e == nil {
		e = echo.New()
	}

	cars := e.Group("cars")
	cars.GET("/:car_id", w.GetCar)
	cars.POST("", w.CreateCar)

	return e
}
