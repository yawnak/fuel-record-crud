package vehicle

import (
	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

type (
	FuelGaugeHistory history.History[*record.FuelGauge, event.FuelGaugeChange]
	OdometerHistory  history.History[*record.Odometer, event.OdometerIncrease]
)

type Vehicle struct {
	cr                  car.Car
	fuelHistoryView     FuelGaugeHistory
	odometerHistoryView OdometerHistory
}

func NewVehicle(cr car.Car, firstFuelRecord *record.FuelGauge, firstOdometerRecord *record.Odometer) (Vehicle, error) {
	if err := cr.Validate(); err != nil {
		return Vehicle{}, err
	}
	vh := Vehicle{
		cr: cr,
	}
	if firstFuelRecord != nil {
		if err := firstFuelRecord.Validate(); err != nil {
			return Vehicle{}, err
		}
		vh.fuelHistoryView = FuelGaugeHistory(history.NewHistory[*record.FuelGauge](firstFuelRecord))
	}

	if firstOdometerRecord != nil {
		if err := firstOdometerRecord.Validate(); err != nil {
			return Vehicle{}, err
		}
		vh.odometerHistoryView = OdometerHistory(history.NewHistory[*record.Odometer](firstOdometerRecord))
	}

	return vh, nil
}

func (vh Vehicle) Car() car.Car {
	return vh.cr
}
