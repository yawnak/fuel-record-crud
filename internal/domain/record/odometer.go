package record

import (
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
)

type Odometer struct {
	event       event.OdometerIncrease
	nextEventId uuid.NullUUID
	prevEventId uuid.NullUUID
}

func NewFirstOdometer(initOdometer float64) (Odometer, error) {
	odom, err := event.NewOdometerIncrease(0, initOdometer)
	return Odometer{
		event: odom,
	}, err
}

func (record Odometer) IsFirst() bool {
	return !record.prevEventId.Valid
}

func (record Odometer) IsLast() bool {
	return !record.nextEventId.Valid
}

func (record *Odometer) NewNext(event event.OdometerIncrease) (*Odometer, error) {
	if record.IsLast() {
		return nil, ErrFuelGaugeRecordNotLast
	}
	return &Odometer{
		event:       event,
		prevEventId: uuid.NullUUID{UUID: record.event.Id(), Valid: true},
	}, nil
}

func (record Odometer) Validate() error {
	return nil
}
