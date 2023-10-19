package car

import (
	"github.com/google/uuid"
)

type (
	FuelHistory     History[*FuelRecord]
	OdometerHistory History[*OdometerRecord]
)
type Car struct {
	id                 uuid.UUID
	make               string
	model              string
	year               int8
	lastFuelRecord     FuelRecord
	lastOdometerRecord OdometerRecord
	//currentOdometerRecord History[*OdometerRecord]
}

func New(make, model string, year int8, currentFuel float64) Car {
	return Car{
		id:    uuid.New(),
		make:  make,
		model: model,
		year:  year,
		//currentFuelRecord: newFuelRecord(currentFuel),
	}
}

func UnmarshalCarFromDatabase(
	id uuid.UUID, make string, model string, year int8, lastFuelRecord FuelRecord,
) Car {
	return Car{
		id:    id,
		make:  make,
		model: model,
		year:  year,
		//currentFuelRecord: lastFuelRecord,
	}
}
