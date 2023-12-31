package event

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNegativeLiters = errors.New("fuel gauge can't have negative liters")
)

type FuelGaugeChange struct {
	id                uuid.UUID
	currentFuelLiters float64
	differenceLiters  float64
	createdAt         time.Time
}

func NewFuelGaugeChange(differenceLiters, currentFuelLiters float64, creationTime time.Time) (FuelGaugeChange, error) {
	if currentFuelLiters < 0 {
		return FuelGaugeChange{}, ErrNegativeLiters
	}
	return FuelGaugeChange{
		id:                uuid.New(),
		differenceLiters:  differenceLiters,
		currentFuelLiters: currentFuelLiters,
		createdAt:         creationTime,
	}, nil
}

func UnmarashalFuelGaugeChangeFromDB(
	id uuid.UUID, currentFuelLiters float64, difference float64, createdAt time.Time) FuelGaugeChange {
	return FuelGaugeChange{
		id:                id,
		currentFuelLiters: currentFuelLiters,
		differenceLiters:  difference,
		createdAt:         createdAt,
	}
}

func (event FuelGaugeChange) Id() uuid.UUID {
	return event.id
}

func (event FuelGaugeChange) CurrentFuelLiters() float64 {
	return event.currentFuelLiters
}

func (event FuelGaugeChange) DifferenceLiters() float64 {
	return event.differenceLiters
}

func (event FuelGaugeChange) CreatedAt() time.Time {
	return event.createdAt
}
