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

func NewOdometerIncrease(differenceKilometers, currentKilometers float64) (OdometerIncrease, error) {
	if differenceKilometers < 0 {
		return OdometerIncrease{}, ErrOdometerNegativeDifference
	}
	return OdometerIncrease{
		id:                   uuid.New(),
		currentKilometers:    currentKilometers,
		differenceKilometers: differenceKilometers,
		createdAt:            time.Now().UTC(),
	}, nil
}
