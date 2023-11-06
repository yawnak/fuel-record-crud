package record

import (
	"time"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

type Odometer struct {
	event       event.OdometerIncrease
	nextEventId uuid.NullUUID
	prevEventId uuid.NullUUID
}

type OdometerHistory struct {
	history.History[*Odometer, event.OdometerIncrease]
}

func NewFirstOdometer(initOdometer float64, creationTime time.Time) (Odometer, error) {
	odom, err := event.NewOdometerIncrease(0, initOdometer, creationTime)
	return Odometer{
		event: odom,
	}, err
}

func UnmarshalOdometerFromRepo(
	event event.OdometerIncrease,
	nextEventId uuid.NullUUID,
	prevEventId uuid.NullUUID,
) Odometer {
	return Odometer{
		event:       event,
		nextEventId: nextEventId,
		prevEventId: prevEventId,
	}
}

func (record *Odometer) IsFirst() bool {
	return !record.prevEventId.Valid
}

func (record *Odometer) IsLast() bool {
	return !record.nextEventId.Valid
}

func (record *Odometer) Copy() *Odometer {
	return &Odometer{
		event:       record.event,
		nextEventId: record.nextEventId,
		prevEventId: record.prevEventId,
	}
}

func (record *Odometer) NewNext(event event.OdometerIncrease) (*Odometer, error) {
	if record.IsLast() {
		return nil, ErrRecordNotLast
	}
	if !record.event.CreatedAt().After(event.CreatedAt()) {
		return nil, ErrNewNextIsEarlierThenLast
	}
	return &Odometer{
		event:       event,
		prevEventId: uuid.NullUUID{UUID: record.event.Id(), Valid: true},
	}, nil
}

func (record Odometer) Validate() error {
	return nil
}

func (record *Odometer) Event() event.OdometerIncrease {
	return record.event
}
