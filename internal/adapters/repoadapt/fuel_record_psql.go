package repoadapt

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

var (
	ErrTxRequired = errors.New("transaction required")
)

func EntFuelRecordToFuelGauge(entFuelRecord *ent.FuelRecord) *record.FuelGauge {
	eve := event.UnmarashalFuelGaugeChangeFromDB(
		entFuelRecord.ID, entFuelRecord.CurrentFuelLiters,
		entFuelRecord.Difference, entFuelRecord.CreatedAt)
	var prevId, nextId uuid.NullUUID
	if entFuelRecord.Edges.Prev != nil {
		prevId = uuid.NullUUID{UUID: entFuelRecord.Edges.Prev.ID, Valid: true}
	}
	if entFuelRecord.NextFuelRecordID != nil {
		nextId = uuid.NullUUID{UUID: *entFuelRecord.NextFuelRecordID, Valid: true}
	}
	return lo.ToPtr(record.UnmarshalFuelGaugeFromRepo(eve, nextId, prevId))
}

type FuelRecordRepoPSQL struct {
	client *ent.FuelRecordClient
}

// if fuelRecord is nil, does nothing
func (repo *FuelRecordRepoPSQL) CreateFuelGaugeRecord(ctx context.Context, carId uuid.UUID, fuelRecord *record.FuelGauge) error {
	if fuelRecord == nil {
		return nil
	}
	_, err := repo.client.Create().
		SetID(fuelRecord.Id()).
		SetCurrentFuelLiters(fuelRecord.Event().CurrentFuelLiters()).
		SetDifference(fuelRecord.Event().DifferenceLiters()).
		SetCreatedAt(fuelRecord.Event().CreatedAt()).
		SetCarID(carId).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FuelRecordRepoPSQL) GetCarFuelRecordsCount(ctx context.Context, carId uuid.UUID) (int, error) {
	return repo.client.Query().Where(fuelrecord.CarIDEQ(carId)).Count(ctx)
}

// Returns ErrCarHasNoFuelGaugeRecords if car has no fuel records
// if error happened during query, returns that error
// else returns nil
func (repo *FuelRecordRepoPSQL) CarHasFuelRecords(ctx context.Context, carId uuid.UUID) error {
	count, err := repo.GetCarFuelRecordsCount(ctx, carId)
	if err != nil {
		return err
	}
	if count == 0 {
		return record.ErrCarHasNoFuelGaugeRecords
	}
	return nil
}

func (repo *FuelRecordRepoPSQL) GetFuelHistory(ctx context.Context, carId uuid.UUID) (record.FuelGaugeHistory, error) {
	entFuelRecords, err := repo.client.Query().Where(fuelrecord.CarIDEQ(carId)).All(ctx)
	if err != nil {
		return record.FuelGaugeHistory{}, err
	}
	records := lo.Map(entFuelRecords, func(entRec *ent.FuelRecord, _ int) *record.FuelGauge {
		return EntFuelRecordToFuelGauge(entRec)
	})
	return record.FuelGaugeHistory{
		History: history.UnmarshalHistory(records),
	}, nil

}
