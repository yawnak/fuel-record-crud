package repoadapt

import (
	"context"
	dbsql "database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
	_ "github.com/yawnak/fuel-record-crud/ent/runtime"
	"github.com/yawnak/fuel-record-crud/internal/services"
)

var (
	testDbName     = "cars"
	testDbUser     = "myuser"
	testDbPassword = "mypassword"
)

func RunPostgresContainer(ctx context.Context, dbName, dbUser, dbPassword string) (*postgres.PostgresContainer, func() error, error) {
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:14-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, nil, err
	}

	// Clean up the container
	cleanup := func() error {
		if err := postgresContainer.Terminate(ctx); err != nil {
			return err
		}
		return nil
	}

	return postgresContainer, cleanup, nil
}

func Open(databaseUrl string) *ent.Client {
	db, err := dbsql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)) // ent.Debug(), ent.Log(log.Println)
}

func TestCRUD(t *testing.T) {
	ctx := context.Background()
	container, cleanup, err := RunPostgresContainer(ctx, testDbName, testDbUser, testDbPassword)
	if err != nil {
		t.Fatal("error creating container:", err)
	}
	defer cleanup()
	connStr, err := container.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	if err != nil {
		t.Fatal("error getting connection string:", err)
	}

	client := Open(connStr)
	defer client.Close()
	client.Schema.Create(ctx)
	//fmt.Print("car creation:\n")
	_, err = client.Car.Create().SetID(uuid.New()).SetMake("Shelby").SetModel("GT500").SetYear(2020).Save(ctx)
	if err != nil {
		t.Fatal("error creating car:", err)
	}

	carUUID := uuid.MustParse("3a19ddf6-58b6-4a18-8965-e2da512ab42f")
	fuelUUID := uuid.MustParse("e8cb5435-adf0-4169-bb44-8a7e2a2d6d21")
	secondFuelUUID := uuid.MustParse("232392e5-fddb-4a17-b968-a815c4649405")
	thirdFuelUUID := uuid.MustParse("2122daec-de4d-4d4b-9231-fac5c4c1d5e8")
	fourthFuelUUID := uuid.MustParse("17422fd9-2d3a-4a05-b69b-5062e3100097")
	// odometerUUID := uuid.MustParse("cd70472f-6eca-492c-a6e4-8861254fc342")
	// secondOdometerUUID := uuid.MustParse("c908b919-ca1a-4f74-b020-bf9bf2bddec9")
	var year int32 = 2010
	client.Car.Create().
		SetID(carUUID).
		SetMake("Toyota").
		SetModel("Corolla").
		SetYear(year).SaveX(ctx)

	f1 := client.FuelRecord.Create().SetID(fuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(10.0).SetDifference(0).SetCreatedAt(time.Now())
	f2 := client.FuelRecord.Create().SetID(secondFuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(20.0).SetDifference(10.0).SetPrevID(fuelUUID).SetCreatedAt(time.Now())
	f3 := client.FuelRecord.Create().SetID(thirdFuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(30.0).SetDifference(10.0).SetPrevID(secondFuelUUID).SetCreatedAt(time.Now())
	f4 := client.FuelRecord.Create().SetID(fourthFuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(40.0).SetDifference(10.0).SetPrevID(thirdFuelUUID).SetCreatedAt(time.Now())

	_, err = client.FuelRecord.CreateBulk(f1, f2, f3, f4).Save(ctx)
	if err != nil {
		t.Fatal("error creating fuel records:", err)
	}
	records, err := client.FuelRecord.Query().Where(fuelrecord.IDIn(fuelUUID, secondFuelUUID, thirdFuelUUID, fourthFuelUUID)).All(ctx)
	if err != nil {
		t.Fatal("error querying fuel records:", err)
	}
	fmt.Println("RECORDS:")
	PrintSlice(records)

	found, _ := lo.Find(records, func(rec *ent.FuelRecord) bool { return rec.ID == fuelUUID })
	foundNext, err := client.FuelRecord.QueryNext(found).All(ctx)
	if err != nil {
		t.Fatal("error querying next fuel record:", err)
	}
	fmt.Println("NEXT RECORDS:")
	PrintSlice(foundNext)

	//fmt.Println(client.Car.QueryFuelRecords(res).Order(fuelrecord.ByCreatedAt(entsql.OrderDesc())))
}

func PrintSlice[T any](s []T) {
	for i, v := range s {
		fmt.Println(i, ":", v)
	}
}

func TestService(t *testing.T) {
	ctx := context.Background()
	container, cleanup, err := RunPostgresContainer(ctx, testDbName, testDbUser, testDbPassword)
	if err != nil {
		t.Fatal("error creating container:", err)
	}
	defer cleanup()
	connStr, err := container.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	if err != nil {
		t.Fatal("error getting connection string:", err)
	}

	client := Open(connStr)
	defer client.Close()
	client.Schema.Create(ctx)
	//fmt.Print("car creation:\n")
	cl := &Client{cl: client}
	serv := services.NewVehicleService(cl)
	serv.CreateVehicle(ctx, "Shelby", "GT500", 2020, nil, time.Now(), nil, time.Now())
	cars, err := client.Car.Query().WithFuelRecords().All(ctx)
	if err != nil {
		t.Fatal("error querying cars:", err)
	}
	fmt.Println("CARS:")
	PrintCars(cars)
}

func PrintCars(cars []*ent.Car) {
	for _, car := range cars {
		fmt.Println("CAR:", car)
		fmt.Println("FUEL RECORDS:", car.Edges.FuelRecords)
	}
}
