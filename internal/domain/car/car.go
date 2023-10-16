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

func (car Car) Id() uuid.UUID {
	return car.id
}

func (car Car) Make() string {
	return car.make
}

func (car Car) Model() string {
	return car.model
}

func (car Car) Year() int64 {
	return car.year
}

func (car Car) LastFuelRecord() FuelRecord {
	return car.lastFuelRecord
}
