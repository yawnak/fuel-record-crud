package car

import "github.com/google/uuid"

type Car struct {
	id             uuid.UUID
	make           string
	model          string
	year           int64
	lastFuelRecord FuelRecord
}

func New(make, model string, year int64, currentFuel float64) Car {
	return Car{
		id:             uuid.New(),
		make:           make,
		model:          model,
		year:           year,
		lastFuelRecord: newFuelRecord(currentFuel),
	}
}
