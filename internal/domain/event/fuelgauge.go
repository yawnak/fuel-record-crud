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

func NewFuelGaugeChange(differenceLiters, currentFuelLiters float64) (FuelGaugeChange, error) {
	if currentFuelLiters < 0 {
		return FuelGaugeChange{}, ErrNegativeLiters
	}
	return FuelGaugeChange{
		id:                uuid.New(),
		currentFuelLiters: currentFuelLiters,
		createdAt:         time.Now().UTC(),
	}, nil
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

func (event FuelGaugeChange) CreatedAt() string {
	return event.createdAt.Format("2006-01-02 15:04:05")
}
