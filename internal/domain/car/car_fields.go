package car

import "github.com/google/uuid"

func (car *Car) Id() uuid.UUID {
	return car.id
}

func (car *Car) Make() string {
	return car.make
}

func (car *Car) Model() string {
	return car.model
}

func (car *Car) Year() int16 {
	return car.year
}

func (car *Car) SetMake(make string) {
	car.make = make
}

func (car *Car) SetModel(model string) {
	car.model = model
}

func (car *Car) SetYear(year int16) {
	car.year = year
}
