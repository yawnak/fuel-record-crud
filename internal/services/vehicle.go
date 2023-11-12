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
	OR record.OdometerRepo,
] interface {
	domain.Tx
	CarRepo() CR
	FuelGaugeChangeRepo() FR
	OdometerIncreaseRepo() OR
}

type VehicleRepo[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
	OR record.OdometerRepo,
	TX VehicleTx[CR, FR, OR],
] interface {
	CarRepo() CR
	BeginTx(context.Context) (TX, error)
}

type VehicleService[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
	OR record.OdometerRepo,
	TX VehicleTx[CR, FR, OR],
] struct {
	repo VehicleRepo[CR, FR, OR, TX]
}

func NewVehicleService[
	CR car.CarRepo,
	FR record.FuelGaugeRepo,
	OR record.OdometerRepo,
	TX VehicleTx[CR, FR, OR],
](vehicleRepo VehicleRepo[CR, FR, OR, TX],
) *VehicleService[CR, FR, OR, TX] {
	// === Function start === //
	return &VehicleService[CR, FR, OR, TX]{repo: vehicleRepo}
}

func (s *VehicleService[CR, FR, OR, TX]) CreateVehicle(
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

	if err := tx.OdometerIncreaseRepo().CreateOdometerRecord(ctx, vh.Car().Id(), vh.OdometerHistory().Head()); err != nil {
		return nil, fmt.Errorf("failed to create odometer record: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &vh, nil
}

func (s *VehicleService[CR, FR, OR, TX]) GetVehicle(ctx context.Context, vehicleId uuid.UUID) (*vehicle.Vehicle, error) {
	var err error
	vh := vehicle.Vehicle{}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	cr, err := tx.CarRepo().GetCar(ctx, vehicleId)
	if err != nil {
		return nil, fmt.Errorf("failed to get car: %w", err)
	}

	fhist, err := tx.FuelGaugeChangeRepo().GetFuelHistory(ctx, vehicleId)
	if err != nil {
		return nil, fmt.Errorf("failed to get fuel history: %w", err)
	}

	ohist, err := tx.OdometerIncreaseRepo().GetOdometerHistory(ctx, vehicleId)
	if err != nil {
		return nil, fmt.Errorf("failed to get odometer history: %w", err)
	}

	vh = vehicle.BuildVehicle(cr, fhist, ohist)

	return &vh, nil
}

func (s *VehicleService[CR, FR, OR, TX]) GetVehiclesCurrent(ctx context.Context) ([]vehicle.Vehicle, error) {
	var err error
	var cars []car.Car
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	if cars, err = tx.CarRepo().QueryCars(ctx); err != nil {
		return nil, fmt.Errorf("failed to query cars: %w", err)
	}
	var fuelRecs map[string]record.FuelGauge
	if fuelRecs, err = tx.FuelGaugeChangeRepo().QueryLastFuelRecords(ctx); err != nil {
		return nil, fmt.Errorf("failed to query fuel records: %w", err)
	}
	var odometerRecs map[string]record.Odometer
	if odometerRecs, err = tx.OdometerIncreaseRepo().QueryLastOdometerRecords(ctx); err != nil {
		return nil, fmt.Errorf("failed to query odometer records: %w", err)
	}

	vehicles := make([]vehicle.Vehicle, 0, len(cars))
	for _, c := range cars {
		fr, ok := fuelRecs[c.Id().String()]
		var fuelRec *record.FuelGauge
		if !ok {
			fuelRec = nil
		} else {
			fuelRec = &fr
		}
		or, ok := odometerRecs[c.Id().String()]
		var odometerRec *record.Odometer
		if !ok {
			odometerRec = nil
		} else {
			odometerRec = &or
		}
		v, err := VehicleFromOneRecord(c, fuelRec, odometerRec)
		if err != nil {
			return nil, fmt.Errorf("failed to build vehicle: %w", err)
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}
