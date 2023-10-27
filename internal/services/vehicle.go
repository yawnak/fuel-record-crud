package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain"
	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/internal/domain/vehicle"
)

type VehicleTx[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
] interface {
	domain.Tx
	CarRepo() CR
	FuelGaugeChangeRepo() FR
}

type VehicleRepo[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
	TX VehicleTx[CR, FR],
] interface {
	CarRepo() CR
	BeginTx(context.Context) (TX, error)
}

type VehicleService[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
	TX VehicleTx[CR, FR],
] struct {
	repo VehicleRepo[CR, FR, TX]
}

func NewVehicleService[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
	TX VehicleTx[CR, FR],
](vehicleRepo VehicleRepo[CR, FR, TX],
) *VehicleService[CR, FR, TX] {
	// === Function start === //
	return &VehicleService[CR, FR, TX]{repo: vehicleRepo}
}

func (s *VehicleService[CR, FR, TX]) CreateVehicle(
	ctx context.Context, model, make string, year int32,
	initFuel *float64, fuelCreationTime time.Time,
	initOdometer *float64, odometerCreationTime time.Time,
) (*vehicle.Vehicle, error) {
	///////////////////////////
	vh, err := vehicle.NewVehicle(
		model, make, year,
		initFuel, fuelCreationTime,
		initOdometer, odometerCreationTime)

	if err != nil {
		return nil, fmt.Errorf("failed to create vehicle: %w", err)
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = tx.FuelGaugeChangeRepo().CarHasFuelRecords(ctx, vh.Car().Id())
	switch {
	case err == nil:
		return nil, record.ErrCarHasFuelGaugeRecords
	case errors.Is(err, record.ErrCarHasNoFuelGaugeRecords):
	default:
		return nil, fmt.Errorf("failed to check if car has fuel records: %w", err)
	}

	if err := tx.CarRepo().CreateCar(ctx, vh.Car()); err != nil {
		return nil, fmt.Errorf("failed to create car: %w", err)
	}

	if err := tx.FuelGaugeChangeRepo().CreateFuelGaugeRecord(ctx, vh.Car().Id(), vh.FuelHistory().Head()); err != nil {
		return nil, fmt.Errorf("failed to create fuel gauge record: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &vh, nil
}

func (s *VehicleService[CR, FR, TX]) GetVehicle(ctx context.Context, carId uuid.UUID) (*vehicle.Vehicle, error) {
	var err error
	vh := vehicle.Vehicle{}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	cr, err := tx.CarRepo().GetCar(ctx, carId)
	if err != nil {
		return nil, fmt.Errorf("failed to get car: %w", err)
	}

	fhist, err := tx.FuelGaugeChangeRepo().GetFuelHistory(ctx, carId)
	if err != nil {
		return nil, fmt.Errorf("failed to get fuel history: %w", err)
	}

	vh = vehicle.BuildVehicle(cr, fhist, record.OdometerHistory{})

	return &vh, nil
}
