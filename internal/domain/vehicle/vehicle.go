package vehicle

import (
	"time"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

type (
	FuelGaugeHistory struct {
		history.History[*record.FuelGauge, event.FuelGaugeChange]
	}
	OdometerHistory struct {
		history.History[*record.Odometer, event.OdometerIncrease]
	}
)

type Vehicle struct {
	cr                  car.Car
	fuelHistoryView     FuelGaugeHistory
	odometerHistoryView OdometerHistory
}

func NewVehicle(
	model, make string, year int32,
	initFuel *float64, fuelCreationTime time.Time,
	initOdometer *float64, odometerCreationTime time.Time,
) (Vehicle, error) {

	vh := Vehicle{}
	cr := car.New(model, make, year)
	vh.cr = cr

	if initFuel != nil {
		rec, err := record.NewFirstFuelGauge(*initFuel, fuelCreationTime)
		if err != nil {
			return Vehicle{}, err
		}
		vh.fuelHistoryView = FuelGaugeHistory{history.NewHistory(&rec)}
	}

	if initOdometer != nil {
		rec, err := record.NewFirstOdometer(*initOdometer, odometerCreationTime)
		if err != nil {
			return Vehicle{}, err
		}
		vh.odometerHistoryView = OdometerHistory{history.NewHistory(&rec)}
	}

	return vh, nil
}

func (vh *Vehicle) Car() *car.Car {
	return &vh.cr
}

func (vh *Vehicle) FuelHistory() *FuelGaugeHistory {
	return &vh.fuelHistoryView
}

func (vh *Vehicle) OdometerHistory() *OdometerHistory {
	return &vh.odometerHistoryView
}
