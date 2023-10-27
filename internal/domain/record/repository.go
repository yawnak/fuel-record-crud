package record

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCarHasNoFuelGaugeRecords = errors.New("car has no fuel gauge records")
	ErrCarHasNoOdometerRecords  = errors.New("car has no odometer records")
	ErrCarHasFuelGaugeRecords   = errors.New("car has fuel gauge records")
	ErrCarHasOdometerRecords    = errors.New("car has odometer records")
)

type FuelGaugeRepo interface {
	// if fuelGauge is nil, does nothing
	CreateFuelGaugeRecord(ctx context.Context, carId uuid.UUID, fuelGauge *FuelGauge) error
	// returns ErrCarHasNoFuelGaugeRecords if car has no fuel records.
	// returns error if error happened during query.
	CarHasFuelRecords(ctx context.Context, carId uuid.UUID) error
}

type OdometerRepo interface {
	CreateOdometerRecord(ctx context.Context, carId uuid.UUID, odometer *Odometer) error
	//GetOdometerRecord(id uuid.UUID) (Odometer, error)
}
