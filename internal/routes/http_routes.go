package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yawnak/fuel-record-crud/internal/app"
	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
	"github.com/yawnak/fuel-record-crud/internal/routes/request"
	"github.com/yawnak/fuel-record-crud/internal/routes/response"
)

type HTTPWrapper struct {
	app *app.Application
}

func NewHTTPWrapper(app *app.Application) (*HTTPWrapper, error) {
	if app == nil {
		return nil, errors.New("app is nil")
	}
	return &HTTPWrapper{app: app}, nil
}

func (w *HTTPWrapper) GetCar(c echo.Context) error {
	carIdParam := c.Param("car_id")
	carId, err := uuid.Parse(carIdParam)
	if err != nil {
		return httperr.NewInvalidQueryParam(
			"car_id param must be valid uuid v4",
			fmt.Errorf("invalid car_id: %s", carIdParam))
	}

	car, err := w.app.Queries.Car.GetHandler.Handle(c.Request().Context(), carId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK,
		response.Car{
			ID:    car.Id().String(),
			Make:  car.Make(),
			Model: car.Model(),
			Year:  car.Year(),
		})
}

func (w *HTTPWrapper) CreateCar(c echo.Context) error {
	req := request.CreateCar{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	//w.app.Commands.Car.Create.Handle(c.Request().Context(), )
	return nil
}
