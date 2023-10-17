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
	return FuelRecord{
		id:               id,
		previousRecordId: previousRecordId,
		nextRecordId:     nextRecordId,
		currentFuel:      currentFuel,
		difference:       difference,
		createdAt:        createdAt}
}

func (record *FuelRecord) NewNext(difference float64) (*FuelRecord, error) {
	if record.nextRecordId.Valid {
		return nil, ErrNextRecordAlreadyExists
	}

	newRecord := FuelRecord{
		id: uuid.New(),
		previousRecordId: uuid.NullUUID{
			Valid: true,
			UUID:  record.id,
		},
		difference:  difference,
		currentFuel: record.currentFuel + difference,
		createdAt:   time.Now().UTC(),
	}

	record.nextRecordId = uuid.NullUUID{
		Valid: true,
		UUID:  newRecord.id,
	}
	return &newRecord, nil
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

func (fuel FuelRecord) IsFirst() bool {
	return fuel.previousRecordId.Valid
}

func (fuel FuelRecord) IsLast() bool {
	return fuel.nextRecordId.Valid
}
