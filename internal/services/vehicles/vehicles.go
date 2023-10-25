package vehicles

import (
	"context"
	"fmt"

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
	FuelGaugeChangeRepo() FR
	BeginTx(context.Context) (TX, error)
}

type VehicleService[CR car.CarRepo,
	FR record.FuelGaugeRepo,
	TX VehicleTx[CR, FR],
] struct {
	repo VehicleRepo[CR, FR, TX]
}

func NewVehicleService[
	CR car.CarRepo, FR record.FuelGaugeRepo, TX VehicleTx[CR, FR],
](vehicleRepo VehicleRepo[CR, FR, TX]) *VehicleService[CR, FR, TX] {
	return &VehicleService[CR, FR, TX]{repo: vehicleRepo}
}

func (s *VehicleService[CR, FR, TX]) CreateVehicle(
	ctx context.Context, model, make string, year int32,
	initFuel *float64,
	initOdometer *float64,
) (vehicle.Vehicle, error) {
	var err error
	newCar := car.New(make, model, year)
	if err := newCar.Validate(); err != nil {
		return vehicle.Vehicle{}, fmt.Errorf("failed to validate car: %w", err)
	}
	var fuelGaugeRecord *record.FuelGauge
	if initFuel != nil {
		var temp record.FuelGauge
		if temp, err = record.NewFirstFuelGauge(*initFuel); err != nil {
			return vehicle.Vehicle{}, fmt.Errorf("failed to create fuel gauge record: %w", err)
		}
		fuelGaugeRecord = &temp
	}

	var odometerRecord *record.Odometer
	if initOdometer != nil {
		var temp record.Odometer
		if temp, err = record.NewFirstOdometer(*initOdometer); err != nil {
			return vehicle.Vehicle{}, fmt.Errorf("failed to create odometer record: %w", err)
		}
		odometerRecord = &temp
	}

	vh, err := vehicle.NewVehicle(newCar, fuelGaugeRecord, odometerRecord)
	if err != nil {
		return vehicle.Vehicle{}, fmt.Errorf("failed to create vehicle: %w", err)
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return vehicle.Vehicle{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	if _, err := tx.CarRepo().CreateCar(ctx, vh.Car()); err != nil {
		return vehicle.Vehicle{}, fmt.Errorf("failed to create car: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return vehicle.Vehicle{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return vehicle.Vehicle{}, nil
}
