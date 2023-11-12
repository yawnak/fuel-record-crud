package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/ent/car"
	_ "github.com/yawnak/fuel-record-crud/ent/runtime"
)

var configPath string = "configs/development.local.yaml"

func main() {
	repo := initRepo()

	ctx := context.Background()
	clearVehicles(ctx, repo)
	seedVehicles(ctx, repo)
}

func clearVehicles(ctx context.Context, client *ent.Client) {
	selectIds := client.Car.Query().Select(car.FieldID).StringsX(ctx)
	ids := lo.Map(selectIds, func(id string, _ int) uuid.UUID {
		return uuid.MustParse(id)
	})
	fmt.Printf("found %d cars in db. Deleting...\n", len(ids))
	for _, u := range ids {
		client.Car.DeleteOneID(u).ExecX(ctx)
	}
	fmt.Println("done deleting.")
}

func seedVehicles(ctx context.Context, client *ent.Client) {
	id := uuid.New()
	client.Car.Create().SetID(id).SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
	client.Car.Create().SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
	client.Car.Create().SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
	client.Car.Create().SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
	client.Car.Create().SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
	client.Car.Create().SetModel("Toyota").SetMake("NO RECORDS").SetYear(2010).SaveX(ctx)
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
