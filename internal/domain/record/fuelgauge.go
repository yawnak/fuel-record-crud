package record

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
)

var (
	ErrRecordNotLast = errors.New("record not last")
)

type FuelGauge struct {
	event       event.FuelGaugeChange
	nextEventId uuid.NullUUID
	prevEventId uuid.NullUUID
}

func NewFirstFuelGauge(initFuel float64, creationTime time.Time) (FuelGauge, error) {
	eve, err := event.NewFuelGaugeChange(0, initFuel, creationTime)
	return FuelGauge{
		event: eve,
	}, err
}

func (record FuelGauge) IsFirst() bool {
	return !record.prevEventId.Valid
}

func (record FuelGauge) IsLast() bool {
	return !record.nextEventId.Valid
}

func (record *FuelGauge) NewNext(event event.FuelGaugeChange) (*FuelGauge, error) {
	if record.IsLast() {
		return nil, ErrRecordNotLast
	}
	if record.event.CreatedAt().After(event.CreatedAt()) {
		return nil, ErrNewNextIsEarlierThenLast
	}

	if record.event.CurrentFuelLiters()+event.DifferenceLiters() != event.CurrentFuelLiters() {
		return nil, ErrRecordStatesDontMatch
	}

	return &FuelGauge{
		event:       event,
		prevEventId: uuid.NullUUID{UUID: record.event.Id(), Valid: true},
	}, nil
}

func (record FuelGauge) Id() uuid.UUID {
	return record.event.Id()
}

func (record FuelGauge) Event() event.FuelGaugeChange {
	return record.event
}

func (record FuelGauge) CreatedAt() time.Time {
	return record.event.CreatedAt()
}

func (record FuelGauge) Validate() error {
	return nil
}
