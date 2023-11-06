package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yawnak/fuel-record-crud/ent"
	_ "github.com/yawnak/fuel-record-crud/ent/runtime"
)

var configPath string = "configs/development.local.yaml"

func main() {
	repo := initRepo()

	ctx := context.Background()
	seedVehicles(ctx, repo)
}

func seedVehicles(ctx context.Context, client *ent.Client) {
	id := uuid.New()
	client.Car.Create().SetID(id).SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
	//mercedes
	id = uuid.New()
	client.Car.Create().SetID(id).SetModel("Mercedes").SetMake("FUEL RECORD").SetYear(2015).SaveX(ctx)
	client.FuelRecord.Create().SetCarID(id).SetCurrentFuelLiters(150.67).SetDifference(0).SetCreatedAt(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	).SaveX(ctx)
	// audi
	id = uuid.New()
	client.Car.Create().SetID(id).SetModel("Audi").SetMake("ODOMETER RECORD").SetYear(2018).SaveX(ctx)
	client.OdometerRecord.Create().SetCarID(id).SetCurrentKilometers(1000.576).SetDifference(0).SetCreatedAt(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	).SaveX(ctx)
	// bmw
	id = uuid.New()
	client.Car.Create().SetID(id).SetModel("BMW").SetMake("FUEL AND ODOMETER RECORDS").SetYear(2019).SaveX(ctx)
	client.FuelRecord.Create().SetCarID(id).SetCurrentFuelLiters(150.67).SetDifference(0).SetCreatedAt(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	).SaveX(ctx)
	client.OdometerRecord.Create().SetCarID(id).SetCurrentKilometers(1000.576).SetDifference(0).SetCreatedAt(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	).SaveX(ctx)
}

func Must[T any](t T, err error) T {
	if err != nil {
		log.Panicln(err)
	}
	return t
}
