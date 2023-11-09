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
	QueryLastFuelRecords(ctx context.Context) (map[string]FuelGauge, error)
	GetFuelHistory(ctx context.Context, carId uuid.UUID) (FuelGaugeHistory, error)
}

type OdometerRepo interface {
	//CreateOdometerRecord(ctx context.Context, carId uuid.UUID, odometer *Odometer) error
	QueryLastOdometerRecords(ctx context.Context) (map[string]Odometer, error)
	GetOdometerHistory(ctx context.Context, vehicleId uuid.UUID) (OdometerHistory, error)
}
