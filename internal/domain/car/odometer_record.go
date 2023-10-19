package car

import (
	"time"

	"github.com/google/uuid"
)

type OdometerRecord struct {
	id                uuid.UUID
	previousRecordId  uuid.NullUUID
	nextRecordId      uuid.NullUUID
	currentKilometers float64
	difference        float64
	createdAt         time.Time
}

func (record *OdometerRecord) newNext(difference float64) (*OdometerRecord, error) {
	if difference < 0 {
		return nil, ErrOdometerNegativeDifference
	}
	if record.nextRecordId.Valid {
		return nil, ErrNextRecordAlreadyExists
	}
	return &OdometerRecord{
		id: uuid.New(),
		previousRecordId: uuid.NullUUID{
			Valid: true,
			UUID:  record.id,
		},
		difference:        difference,
		currentKilometers: record.currentKilometers + difference,
		createdAt:         time.Now().UTC(),
	}, nil
}

func (record OdometerRecord) Id() uuid.UUID {
	return record.id
}

func (record OdometerRecord) PreviousRecordId() uuid.NullUUID {
	return record.previousRecordId
}

func (record OdometerRecord) NextRecordId() uuid.NullUUID {
	return record.nextRecordId
}

func (record OdometerRecord) CurrentKilometers() float64 {
	return record.currentKilometers
}

func (record OdometerRecord) Difference() float64 {
	return record.difference
}

func (record OdometerRecord) CreatedAt() time.Time {
	return record.createdAt
}

func (record OdometerRecord) IsFirst() bool {
	return record.previousRecordId.Valid
}

func (record OdometerRecord) IsLast() bool {
	return record.nextRecordId.Valid
}
