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

func (car *Car) Year() int8 {
	return car.year
}

// func (car *Car) CurrentFuelRecord() FuelRecord {
// 	return
// }

// func (car *Car) CurrentOdometerRecord() OdometerRecord {
// 	return car.currentOdometerRecord
// }

func (car *Car) SetMake(make string) {
	car.make = make
}

func (car *Car) SetModel(model string) {
	car.model = model
}

func (car *Car) SetYear(year int8) {
	car.year = year
}
