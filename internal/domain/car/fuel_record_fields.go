package car

import (
	"time"

	"github.com/google/uuid"
)

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
