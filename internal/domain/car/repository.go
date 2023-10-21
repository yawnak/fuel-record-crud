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
