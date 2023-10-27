package record

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

var (
	ErrRecordNotLast = errors.New("record not last")
)

type FuelGauge struct {
	event       event.FuelGaugeChange
	nextEventId uuid.NullUUID
	prevEventId uuid.NullUUID
}

type FuelGaugeHistory struct {
	history.History[*FuelGauge, event.FuelGaugeChange]
}

func NewFirstFuelGauge(initFuel float64, creationTime time.Time) (FuelGauge, error) {
	eve, err := event.NewFuelGaugeChange(0, initFuel, creationTime)
	return FuelGauge{
		event: eve,
	}, err
}

func UnmarshalFuelGaugeFromRepo(
	event event.FuelGaugeChange,
	nextEventId uuid.NullUUID,
	prevEventId uuid.NullUUID,
) FuelGauge {
	return FuelGauge{
		event:       event,
		nextEventId: nextEventId,
		prevEventId: prevEventId,
	}
}

func (record *FuelGauge) IsFirst() bool {
	return !record.prevEventId.Valid
}

func (record *FuelGauge) IsLast() bool {
	return !record.nextEventId.Valid
}

func (record *FuelGauge) Copy() *FuelGauge {
	return &FuelGauge{
		event:       record.event,
		nextEventId: record.nextEventId,
		prevEventId: record.prevEventId,
	}
}

func (record *FuelGauge) NewNext(event event.FuelGaugeChange) (*FuelGauge, error) {
	if !record.IsLast() {
		return nil, ErrRecordNotLast
	}
	if record.event.CreatedAt().After(event.CreatedAt()) {
		return nil, ErrNewNextIsEarlierThenLast
	}

	if record.event.CurrentFuelLiters()+event.DifferenceLiters() != event.CurrentFuelLiters() {
		return nil, ErrRecordStatesDontMatch
	}

	record.nextEventId = uuid.NullUUID{UUID: event.Id(), Valid: true}

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

func (record FuelGauge) NextEventId() uuid.NullUUID {
	return record.nextEventId
}

func (record FuelGauge) PrevEventId() uuid.NullUUID {
	return record.prevEventId
}
