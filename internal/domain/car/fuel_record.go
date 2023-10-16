package car

import (
	"time"

	"github.com/google/uuid"
)

type FuelRecord struct {
	id               uuid.UUID
	previousRecordId uuid.NullUUID
	nextRecordId     uuid.NullUUID
	currentFuel      float64
	difference       float64
	createdAt        time.Time
}

func newFuelRecord(currentFuel float64) FuelRecord {
	return FuelRecord{
		id:          uuid.New(),
		currentFuel: currentFuel,
		createdAt:   time.Now().UTC(),
	}
}

func (fuel FuelRecord) NewNext(difference float64) (FuelRecord, error) {
	var zero FuelRecord
	if fuel.nextRecordId.Valid {
		return zero, ErrNextRecordAlreadyExists
	}
	return FuelRecord{
		id: uuid.New(),
		previousRecordId: uuid.NullUUID{
			Valid: true,
			UUID:  uuid.New(),
		},
		difference:  difference,
		currentFuel: fuel.currentFuel + difference,
		createdAt:   time.Now().UTC(),
	}, nil
}
