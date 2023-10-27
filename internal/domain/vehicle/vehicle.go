package vehicle

import (
	"time"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

type Vehicle struct {
	cr                  car.Car
	fuelHistoryView     record.FuelGaugeHistory
	odometerHistoryView record.OdometerHistory
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
		vh.fuelHistoryView = record.FuelGaugeHistory{History: history.NewHistory(&rec)}
	}

	if initOdometer != nil {
		rec, err := record.NewFirstOdometer(*initOdometer, odometerCreationTime)
		if err != nil {
			return Vehicle{}, err
		}
		vh.odometerHistoryView = record.OdometerHistory{History: history.NewHistory(&rec)}
	}

	return vh, nil
}

func BuildVehicle(
	cr car.Car,
	fuelHistoryView record.FuelGaugeHistory,
	odometerHistoryView record.OdometerHistory,
) Vehicle {
	return Vehicle{
		cr:                  cr,
		fuelHistoryView:     fuelHistoryView,
		odometerHistoryView: odometerHistoryView,
	}
}

func (vh *Vehicle) Car() *car.Car {
	return &vh.cr
}

func (vh *Vehicle) FuelHistory() *record.FuelGaugeHistory {
	return &vh.fuelHistoryView
}

func (vh *Vehicle) OdometerHistory() *record.OdometerHistory {
	return &vh.odometerHistoryView
}
