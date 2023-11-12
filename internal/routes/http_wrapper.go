package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/internal/app"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/internal/routes/request"
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

func (w *HTTPWrapper) GetVehicle(c echo.Context) error {
	ctx := c.Request().Context()
	vehicleId, err := uuid.Parse(c.Param("vehicle_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	vehicle, err := w.app.Queries.Vehicle.GetVehicle.Handle(ctx, vehicleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_ = vehicle

	type fuelRecord struct {
		CurrentFuel float64
		Difference  float64
		CreatedAt   string
	}

	type odometerRecord struct {
		CurrentOdometer float64
		Difference      float64
		CreatedAt       string
	}

	data := struct {
		FuelRecords     []fuelRecord
		OdometerRecords []odometerRecord
	}{
		FuelRecords: lo.Map(vehicle.FuelHistory().List(), func(rec *record.FuelGauge, _ int) fuelRecord {
			return fuelRecord{
				CurrentFuel: rec.Event().CurrentFuelLiters(),
				Difference:  rec.Event().DifferenceLiters(),
				CreatedAt:   rec.Event().CreatedAt().Format(time.DateTime),
			}
		}),
		OdometerRecords: lo.Map(vehicle.OdometerHistory().List(), func(rec *record.Odometer, _ int) odometerRecord {
			return odometerRecord{
				CurrentOdometer: rec.Event().CurrentKilometers(),
				Difference:      rec.Event().DifferenceKilometers(),
				CreatedAt:       rec.Event().CreatedAt().Format(time.DateTime),
			}
		}),
	}

	return c.Render(http.StatusOK, "get_vehicles_resp.htm", data)
}

func (w *HTTPWrapper) CreateVehicle(c echo.Context) error {
	ctx := c.Request().Context()
	request := request.CreateVehicle{}
	err := c.Bind(&request)
	if err != nil {
		fmt.Println("ERROR:", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("REQUEST:", "CAR:", request.Model, request.Make, request.Year,
		"\nFUEL:", lo.FromPtr(request.CurrentFuel), lo.FromPtr(request.FuelCreationTime),
		"\nODOMETER:", lo.FromPtr(request.CurrentOdometer), lo.FromPtr(request.OdometerCreationTime))
	vehicle, err := w.app.Commands.Vehicle.Create.Handle(ctx, request.ToCmd())
	if err != nil {
		fmt.Println("ERROR:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmtFloatPtr := func(f *float64) string {
		if f == nil {
			return "N/A"
		}
		return strconv.FormatFloat(*f, 'f', 3, 64)
	}
	var curFuel *float64
	var curOdometer *float64
	if h := vehicle.FuelHistory().Head(); h != nil {
		curFuel = lo.ToPtr(h.Event().CurrentFuelLiters())
	}
	if h := vehicle.OdometerHistory().Head(); h != nil {
		curOdometer = lo.ToPtr(h.Event().CurrentKilometers())
	}

	data := struct {
		ID              string
		Make            string
		Model           string
		Year            int
		CurrentFuel     string
		CurrentOdometer string
	}{
		ID:              vehicle.Car().Id().String(),
		Make:            vehicle.Car().Make(),
		Model:           vehicle.Car().Model(),
		Year:            int(vehicle.Car().Year()),
		CurrentFuel:     fmtFloatPtr(curFuel),
		CurrentOdometer: fmtFloatPtr(curOdometer),
	}

	return c.Render(http.StatusOK, "create_vehicle_resp.htm", data)
}

func (w *HTTPWrapper) GetFormCreateFuel(c echo.Context) error {
	return c.Render(http.StatusOK, "form_create_fuel.htm", nil)
}

func (w *HTTPWrapper) GetFormCreateOdometer(c echo.Context) error {
	return c.Render(http.StatusOK, "form_create_odometer.htm", nil)
}
