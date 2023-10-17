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

func (record *OdometerRecord) NewNext(difference float64) (OdometerRecord, error) {
	var zero OdometerRecord
	if difference < 0 {
		return zero, ErrOdometerNegativeDifference
	}
	if record.nextRecordId.Valid {
		return zero, ErrNextRecordAlreadyExists
	}
	return OdometerRecord{
		id: uuid.New(),
		previousRecordId: uuid.NullUUID{
			Valid: true,
			UUID:  uuid.New(),
		},
		difference:        difference,
		currentKilometers: record.currentKilometers + difference,
		createdAt:         time.Now().UTC(),
	}, nil
}
