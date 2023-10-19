package repoadapt

import (
	"time"

	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/internal/domain/car"
)

// type carModel struct {
// 	CarId            string `db:"car_id"`
// 	Make             string `db:"make"`
// 	Model            string `db:"model"`
// 	Year             int64  `db:"year"`
// 	LastFuelRecordId string `db:"last_fuel_record_id"`
// }

// type fuelRecordModel struct {
// 	FuelRecordId     string    `db:"fuel_record_id"`
// 	PreviousRecordId *string   `db:"previous_record_id"`
// 	NextRecordId     *string   `db:"next_record_id"`
// 	CurrentFuel      float64   `db:"current_fuel"`
// 	Difference       float64   `db:"difference"`
// 	CreatedAt        time.Time `db:"created_at"`
// }

// func unmarshalCar(car car.Car) carModel {
// 	return carModel{
// 		CarId: car.Id().String(),
// 		Make:  car.Make(),
// 		Model: car.Model(),
// 		Year:  car.Year(),
// 		//LastFuelRecordId: car.CurrentFuelRecord().Id().String(),
// 	}
// }

// func marshalCar(carMd carModel, fuelRecordMd fuelRecordModel) (car.Car, error) {
// 	carid, err := uuid.Parse(carMd.CarId)
// 	if err != nil {
// 		log.Println("error parsing car id", err)
// 	}

// 	if carMd.LastFuelRecordId != fuelRecordMd.FuelRecordId {
// 		log.Println("last fuel record id does not match fuel record id")
// 	}

// 	fuelRecordId, err := uuid.Parse(fuelRecordMd.FuelRecordId)
// 	if err != nil {
// 		log.Println("error parsing last fuel record id", err)
// 	}

// 	previoudRecordId := uuid.NullUUID{}
// 	if fuelRecordMd.PreviousRecordId != nil {
// 		previoudRecordId.Valid = true
// 		previoudRecordId.UUID, err = uuid.Parse(*fuelRecordMd.PreviousRecordId)
// 		if err != nil {
// 			log.Println("error parsing previous record id", err)
// 		}
// 	}
// 	nextRecordId := uuid.NullUUID{}
// 	if fuelRecordMd.NextRecordId != nil {
// 		nextRecordId.Valid = true
// 		nextRecordId.UUID, err = uuid.Parse(*fuelRecordMd.NextRecordId)
// 		if err != nil {
// 			log.Println("error parsing next record id", err)
// 		}
// 	}

// 	return car.UnmarshalCarFromDatabase(
// 		carid,
// 		carMd.Make,
// 		carMd.Model,
// 		carMd.Year,
// 		car.UnmarshalFuelRecordFromDatabase(
// 			fuelRecordId,
// 			previoudRecordId,
// 			nextRecordId,
// 			fuelRecordMd.CurrentFuel,
// 			fuelRecordMd.Difference,
// 			fuelRecordMd.CreatedAt,
// 		),
// 	), nil
// }

type CarRepositoryPSQL struct {
	client *ent.Client
}

func (repo *CarRepositoryPSQL) CreateCar(c car.Car) error {
	repo.client.FuelRecord.Create().
		SetCurrentFuelLiters(0).
		SetDifference(0).
		SetCreatedAt(time.Now())
	repo.client.Car.Create().
		SetID(c.Id()).
		SetMake(c.Make()).
		SetModel(c.Model()).
		SetYear(c.Year())
}
