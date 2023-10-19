package car

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestFuelRecordNewNext(t *testing.T) {
	startUUID := uuid.MustParse("299b34a2-b519-4a3c-800f-1ea06b7c1563")

	cases := []struct {
		name        string
		start       FuelRecord
		difference  float64
		expectStart FuelRecord
		expectNext  FuelRecord
		expectErr   error
	}{
		{
			name: "positive difference",
			start: FuelRecord{
				id:           startUUID,
				nextRecordId: uuid.NullUUID{Valid: false},
				currentFuel:  10,
				difference:   3,
				createdAt:    time.Now().UTC(),
			},
			difference: 5,
			expectStart: FuelRecord{
				id:          startUUID,
				currentFuel: 10,
				difference:  3,
			},
			expectNext: FuelRecord{
				previousRecordId: uuid.NullUUID{Valid: true, UUID: startUUID},
				nextRecordId:     uuid.NullUUID{Valid: false},
				currentFuel:      15,
				difference:       5,
			},
			expectErr: nil,
		},
		{
			name: "negative difference",
			start: FuelRecord{
				id:           startUUID,
				nextRecordId: uuid.NullUUID{Valid: false},
				currentFuel:  10,
				difference:   3,
				createdAt:    time.Now().UTC(),
			},
			difference: -5,
			expectStart: FuelRecord{
				id:          startUUID,
				currentFuel: 10,
				difference:  3,
			},
			expectNext: FuelRecord{
				previousRecordId: uuid.NullUUID{Valid: true, UUID: startUUID},
				nextRecordId:     uuid.NullUUID{Valid: false},
				currentFuel:      5,
				difference:       -5,
			},
			expectErr: nil,
		},
		{
			name: "next record already exists",
			start: FuelRecord{
				id:           startUUID,
				nextRecordId: uuid.NullUUID{Valid: true, UUID: uuid.New()},
				currentFuel:  10,
				difference:   3,
				createdAt:    time.Now().UTC(),
			},
			difference: 5,
			expectErr:  ErrNextRecordAlreadyExists,
		},
	}

	assertFn := func(rec, assert FuelRecord) error {
		if rec.previousRecordId != assert.previousRecordId {
			return fmt.Errorf("expected previousRecordId %v, got %v", assert.previousRecordId, rec.previousRecordId)
		}
		if rec.nextRecordId != assert.nextRecordId {
			return fmt.Errorf("expected nextRecordId %v, got %v", assert.nextRecordId, rec.nextRecordId)
		}
		if rec.currentFuel != assert.currentFuel {
			return fmt.Errorf("expected currentFuel %v, got %v", assert.currentFuel, rec.currentFuel)
		}
		if rec.difference != assert.difference {
			return fmt.Errorf("expected difference %v, got %v", assert.difference, rec.difference)
		}
		return nil
	}

	for i, v := range cases {
		rec, err := v.start.newNext(v.difference)
		if err != v.expectErr {
			t.Errorf("case %d: expected error %v, got %v", i, v.expectErr, err)
		}
		if err != nil {
			continue
		}
		v.expectStart.nextRecordId = uuid.NullUUID{Valid: true, UUID: rec.id}
		if err := assertFn(*rec, v.expectNext); err != nil {
			t.Errorf("case %d expectNext: %v", i, err)
		}
		if err := assertFn(v.start, v.expectStart); err != nil {
			t.Errorf("case %d expectStart: %v", i, err)
		}
	}
}
