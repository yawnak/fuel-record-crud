package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/internal/app"
)

type HTTPWrapper struct {
	app *app.Application
}

func NewHTTPWrapper(app *app.Application) *HTTPWrapper {
	return &HTTPWrapper{app: app}
}

func (w *HTTPWrapper) GetVehiclesPage(c echo.Context) error {
	ctx := c.Request().Context()

	vehicles, err := w.app.Queries.Vehicle.GetVehiclesCurrentData.Handle(ctx)
	if err != nil {
		return err
	}

	fmt.Println("LEN VEHICLES:", len(vehicles))

	type tmplVehicle struct {
		ID              string
		Make            string
		Model           string
		Year            int32
		CurrentFuel     string
		CurrentOdometer string
	}

	fmtFloatPtr := func(f *float64) string {
		if f == nil {
			return "N/A"
		}
		return strconv.FormatFloat(*f, 'f', 3, 64)
	}

	tmplVehicles := make([]tmplVehicle, 0, len(vehicles))
	for _, v := range vehicles {
		var curFuel *float64
		var curOdometer *float64
		if h := v.FuelHistory().Head(); h != nil {
			curFuel = lo.ToPtr(h.Event().CurrentFuelLiters())
		}
		if h := v.OdometerHistory().Head(); h != nil {
			curOdometer = lo.ToPtr(h.Event().CurrentKilometers())
		}

		tmplVehicles = append(tmplVehicles, tmplVehicle{
			ID:              v.Car().Id().String(),
			Make:            v.Car().Make(),
			Model:           v.Car().Model(),
			Year:            v.Car().Year(),
			CurrentFuel:     fmtFloatPtr(curFuel),
			CurrentOdometer: fmtFloatPtr(curOdometer),
		})
	}
	data := struct {
		Vehicles []tmplVehicle
	}{
		Vehicles: tmplVehicles,
	}

	return c.Render(http.StatusOK, "vehicles.htm", data)
}

// func (w *HTTPWrapper) GetCar(c echo.Context) error {
// 	carIdParam := c.Param("car_id")
// 	carId, err := uuid.Parse(carIdParam)
// 	if err != nil {
// 		return httperr.NewInvalidQueryParam(
// 			"car_id param must be valid uuid v4",
// 			fmt.Errorf("invalid car_id: %s", carIdParam))
// 	}

// 	car, err := w.app.Queries.Car.GetHandler.Handle(c.Request().Context(), carId)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK,
// 		response.Car{
// 			ID:    car.Id().String(),
// 			Make:  car.Make(),
// 			Model: car.Model(),
// 			Year:  car.Year(),
// 		})
// }

func (w *HTTPWrapper) GetVehicle(c echo.Context) error {
	// vehicleId, err := uuid.Parse(c.Param("vehicle_id"))
	// if err != nil {
	// 	return c.Render()
	// }

	return nil
}
