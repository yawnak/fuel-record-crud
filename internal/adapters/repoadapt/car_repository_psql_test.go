package repoadapt

import (
	"context"
	dbsql "database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yawnak/fuel-record-crud/ent"
)

func Open(databaseUrl string) *ent.Client {
	db, err := dbsql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func TestCRUD(t *testing.T) {
	client := Open("postgresql://myuser:mypassword@localhost:5432/cars")
	defer client.Close()

	ctx := context.Background()

	client.FuelRecord.Delete().ExecX(ctx)
	client.OdometerRecord.Delete().ExecX(ctx)
	client.Car.Delete().ExecX(ctx)

	carUUID := uuid.MustParse("3a19ddf6-58b6-4a18-8965-e2da512ab42f")
	fuelUUID := uuid.MustParse("e8cb5435-adf0-4169-bb44-8a7e2a2d6d21")
	nextFuelUUID := uuid.MustParse("232392e5-fddb-4a17-b968-a815c4649405")
	odometerUUID := uuid.MustParse("cd70472f-6eca-492c-a6e4-8861254fc342")
	nextOdometerUUID := uuid.MustParse("c908b919-ca1a-4f74-b020-bf9bf2bddec9")
	var year int32 = 2010
	client.Car.Create().
		SetID(carUUID).
		SetMake("Toyota").
		SetModel("Corolla").
		SetYear(year).SaveX(ctx)

	res := client.Car.GetX(ctx, carUUID)
	fmt.Println(res.String())

	client.Car.UpdateOneID(carUUID).SetMake("Mercedes").SaveX(ctx)

	res = client.Car.GetX(ctx, carUUID)
	fmt.Println(res.String())
	firstFuel := &ent.FuelRecord{}
	secondFuel := &ent.FuelRecord{}
	var err error
	if firstFuel, err = client.FuelRecord.Create().SetID(fuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(10).SetDifference(10).SetCreatedAt(time.Now()).Save(ctx); err != nil {
		log.Fatal(err)
	}
	if secondFuel, err = client.FuelRecord.Create().SetID(nextFuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(20).SetDifference(10).SetCreatedAt(time.Now()).
		SetPrev(firstFuel).Save(ctx); err != nil {
		log.Fatal(err)
	}

	client.OdometerRecord.Create().SetID(odometerUUID).SetCarID(carUUID).SetCurrentFuelLiters(100).SetDifference(100).SetCreatedAt(time.Now()).SaveX(ctx)
	client.OdometerRecord.Create().SetID(nextOdometerUUID).SetCarID(carUUID).SetCurrentFuelLiters(200).SetDifference(100).SetCreatedAt(time.Now()).SetPrevID(odometerUUID).SaveX(ctx)

	fmt.Println(client.FuelRecord.QueryNext(firstFuel).All(ctx))
	fmt.Println(client.FuelRecord.QueryPrev(secondFuel).All(ctx))

	//fmt.Println(client.Car.QueryFuelRecords(res).Order(fuelrecord.ByCreatedAt(entsql.OrderDesc())))
}
