package car

import (
	"context"

	"github.com/google/uuid"
)

type CarRepo interface {
	CreateCar(ctx context.Context, car Car) (Car, error)
	GetCar(ctx context.Context, id uuid.UUID) (Car, error)
	QueryCars(ctx context.Context) ([]Car, error)
	//Update(ctx context.Context, updateFunc func(*Car)) (Car, error)
	//Delete(ctx context.Context, id uuid.UUID) error
}
