package event

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrOdometerNegativeDifference = errors.New("odometer negative difference")
)

type OdometerIncrease struct {
	id                   uuid.UUID
	currentKilometers    float64
	differenceKilometers float64
	createdAt            time.Time
}

func NewOdometerIncrease(differenceKilometers, currentKilometers float64, creationTime time.Time) (OdometerIncrease, error) {
	if differenceKilometers < 0 {
		return OdometerIncrease{}, ErrOdometerNegativeDifference
	}
	return OdometerIncrease{
		id:                   uuid.New(),
		currentKilometers:    currentKilometers,
		differenceKilometers: differenceKilometers,
		createdAt:            creationTime,
	}, nil
}

func (event OdometerIncrease) Id() uuid.UUID {
	return event.id
}

func (event OdometerIncrease) CurrentKilometers() float64 {
	return event.currentKilometers
}

func (event OdometerIncrease) DifferenceKilometers() float64 {
	return event.differenceKilometers
}

func (event OdometerIncrease) CreatedAt() time.Time {
	return event.createdAt
}
