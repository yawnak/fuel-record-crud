package repoadapt

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

func EntOdometerRecordToOdometer(entOdometer *ent.OdometerRecord) *record.Odometer {
	eve := event.UnmarshalOdometerIncreaseFromDB(
		entOdometer.ID, entOdometer.CurrentKilometers,
		entOdometer.Difference, entOdometer.CreatedAt)
	var prevId, nextId uuid.NullUUID
	if entOdometer.Edges.Prev != nil {
		prevId = uuid.NullUUID{UUID: entOdometer.Edges.Prev.ID, Valid: true}
	}
	if entOdometer.NextOdometerRecordID != nil {
		nextId = uuid.NullUUID{UUID: *entOdometer.NextOdometerRecordID, Valid: true}
	}
	return lo.ToPtr(record.UnmarshalOdometerFromRepo(eve, nextId, prevId))
}

type OdometerRecordRepoPSQL struct {
	client *ent.OdometerRecordClient
}

func (repo *OdometerRecordRepoPSQL) Create(ctx context.Context, odometerRecord event.OdometerIncrease) (event.OdometerIncrease, error) {
	return event.OdometerIncrease{}, nil
}

func (repo *OdometerRecordRepoPSQL) QueryLastOdometerRecords(ctx context.Context) (map[string]record.Odometer, error) {
	fuelrecords, err := repo.client.Query().Where(odometerrecord.NextOdometerRecordIDIsNil()).All(ctx)
	if err != nil {
		return nil, err
	}
	res := map[string]record.Odometer{}
	for i, _ := range fuelrecords {
		res[fuelrecords[i].CarID.String()] = lo.FromPtr(EntOdometerRecordToOdometer(fuelrecords[i]))
	}
	return res, nil
}

func (repo *OdometerRecordRepoPSQL) GetOdometerHistory(ctx context.Context, carId uuid.UUID) (record.OdometerHistory, error) {
	records, err := repo.client.Query().Where(odometerrecord.CarIDEQ(carId)).All(ctx)
	if err != nil {
		return record.OdometerHistory{}, err
	}
	res := make([]*record.Odometer, len(records))
	for i, _ := range records {
		res[i] = EntOdometerRecordToOdometer(records[i])
	}
	return record.OdometerHistory{
		History: history.UnmarshalHistory(res),
	}, nil
}

func (repo *OdometerRecordRepoPSQL) CreateOdometerRecord(ctx context.Context, carId uuid.UUID, odometer *record.Odometer) error {
	if odometer == nil {
		return nil
	}
	_, err := repo.client.Create().
		SetID(odometer.Event().Id()).
		SetCurrentKilometers(odometer.Event().CurrentKilometers()).
		SetDifference(odometer.Event().DifferenceKilometers()).
		SetCreatedAt(odometer.Event().CreatedAt()).
		SetCarID(carId).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
