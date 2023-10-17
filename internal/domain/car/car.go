package car

import (
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

type (
	FuelHistory     history.History[*FuelRecord]
	OdometerHistory history.History[*OdometerRecord]
)
type Car struct {
	id              uuid.UUID
	make            string
	model           string
	year            int64
	fuelHistory     FuelHistory
	odometerHistory OdometerHistory
	//currentOdometerRecord History[*OdometerRecord]
}

func New(make, model string, year int64, currentFuel float64) Car {
	return Car{
		id:    uuid.New(),
		make:  make,
		model: model,
		year:  year,
		//currentFuelRecord: newFuelRecord(currentFuel),
	}
}

func UnmarshalCarFromDatabase(
	id uuid.UUID, make string, model string, year int64, lastFuelRecord FuelRecord,
) Car {
	return Car{
		id:    id,
		make:  make,
		model: model,
		year:  year,
		//currentFuelRecord: lastFuelRecord,
	}
}

func (car *Car) Id() uuid.UUID {
	return car.id
}

func (car *Car) Make() string {
	return car.make
}

func (car *Car) Model() string {
	return car.model
}

func (car *Car) Year() int64 {
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

func (car *Car) SetYear(year int64) {
	car.year = year
}
