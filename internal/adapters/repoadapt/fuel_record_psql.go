package repoadapt

import (
	"context"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
)

type FuelRecordRepoPSQL struct {
	client *ent.FuelRecordClient
}

func (repo *FuelRecordRepoPSQL) CreateFuelGaugeRecord(ctx context.Context, carId uuid.UUID, fuelRecord record.FuelGauge) (record.FuelGauge, error) {
	return record.FuelGauge{}, nil
}
