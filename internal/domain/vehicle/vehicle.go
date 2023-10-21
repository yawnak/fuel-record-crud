package vehicle

import (
	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/view"
)

type Vehicle struct {
	cr              car.Car
	fuelHistoryView view.FuelGauge
}
