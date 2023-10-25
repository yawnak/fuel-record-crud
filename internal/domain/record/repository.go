package record

import (
	"context"

	"github.com/google/uuid"
)

type FuelGaugeRepo interface {
	CreateFuelGaugeRecord(ctx context.Context, carId uuid.UUID, fuelGauge FuelGauge) (FuelGauge, error)
	//GetFuelGaugeRecord(id uuid.UUID) (FuelGauge, error)
}

type OdometerRepo interface {
	CreateOdometerRecord(ctx context.Context, carId uuid.UUID, odometer Odometer) (Odometer, error)
	//GetOdometerRecord(id uuid.UUID) (Odometer, error)
}
