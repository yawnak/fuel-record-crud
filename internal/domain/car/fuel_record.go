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

// this method creates a first record without prev or next.
// Should be used only when creating a new car.
func newFuelRecord(currentFuel float64) FuelRecord {
	return FuelRecord{
		id:          uuid.New(),
		currentFuel: currentFuel,
		createdAt:   time.Now().UTC(),
	}
}

// this method should only be used when unmarshalling from database
// can create an invalid record
func UnmarshalFuelRecordFromDatabase(
	id uuid.UUID, previousRecordId uuid.NullUUID, nextRecordId uuid.NullUUID,
	currentFuel float64, difference float64, createdAt time.Time,
) FuelRecord {
	return FuelRecord{}
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

func (fuel FuelRecord) Id() uuid.UUID {
	return fuel.id
}

func (fuel FuelRecord) PreviousRecordId() uuid.NullUUID {
	return fuel.previousRecordId
}

func (fuel FuelRecord) NextRecordId() uuid.NullUUID {
	return fuel.nextRecordId
}

func (fuel FuelRecord) CurrentFuel() float64 {
	return fuel.currentFuel
}

func (fuel FuelRecord) Difference() float64 {
	return fuel.difference
}

func (fuel FuelRecord) CreatedAt() time.Time {
	return fuel.createdAt
}
