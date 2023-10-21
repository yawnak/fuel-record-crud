package event

import (
	"context"

	"github.com/google/uuid"
)

type FuelGaugeChangeRepository interface {
	Create(ctx context.Context, fuelgauge FuelGaugeChange) (FuelGaugeChange, error)
	Get(ctx context.Context, id uuid.UUID) (FuelGaugeChange, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
