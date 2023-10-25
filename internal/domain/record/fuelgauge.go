package record

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
)

var (
	ErrFuelGaugeRecordNotLast = errors.New("fuel gauge record not last")
)

type FuelGauge struct {
	event       event.FuelGaugeChange
	nextEventId uuid.NullUUID
	prevEventId uuid.NullUUID
}

func NewFirstFuelGauge(initFuel float64) (FuelGauge, error) {
	eve, err := event.NewFuelGaugeChange(0, initFuel)
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
		return nil, ErrFuelGaugeRecordNotLast
	}
	return &FuelGauge{
		event:       event,
		prevEventId: uuid.NullUUID{UUID: record.event.Id(), Valid: true},
	}, nil
}

func (record FuelGauge) Validate() error {
	return nil
}
