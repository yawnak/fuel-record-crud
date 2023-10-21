package view

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

var (
	ErrFuelGaugeRecordNotLast = errors.New("fuel gauge record not last")
)

type FuelGaugeRecord struct {
	event       event.FuelGaugeChange
	nextEventId uuid.NullUUID
	prevEventId uuid.NullUUID
}

func (record FuelGaugeRecord) IsFirst() bool {
	return !record.prevEventId.Valid
}

func (record FuelGaugeRecord) IsLast() bool {
	return !record.nextEventId.Valid
}

func (record *FuelGaugeRecord) NewNext() (*FuelGaugeRecord, error) {
	if record.IsLast() {
		return nil, ErrFuelGaugeRecordNotLast
	}
	return &FuelGaugeRecord{
		event:       event.FuelGaugeChange{},
		prevEventId: uuid.NullUUID{UUID: record.event.Id(), Valid: true},
	}, nil
}

type FuelGauge struct {
	historyView history.History[*FuelGaugeRecord]
}
