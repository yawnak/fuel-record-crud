package car

import (
	"context"

	"github.com/google/uuid"
)

type CarRepository interface {
	Create(ctx context.Context, car Car) (Car, error)
	Get(ctx context.Context, id uuid.UUID) (Car, error)
	GetAll(ctx context.Context) ([]Car, error)
	Update(ctx context.Context, updateFunc func(*Car)) (Car, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type FuelRecordRepository interface {
	Get(fuelRecordId uuid.UUID) (*FuelRecord, error)
	GetLastForCar(carId uuid.UUID) (*FuelRecord, error)
}

type OdometerRecordRepository interface {
	Get(odometerRecordId uuid.UUID) (*OdometerRecord, error)
	GetLastForCar(carId uuid.UUID) (*OdometerRecord, error)
}

type FuelHistoryRepository interface {
	GetByCarId(carId uuid.UUID) (*FuelHistory, error)
	Update(start *FuelRecord, fuelHistory *FuelHistory) error
}
