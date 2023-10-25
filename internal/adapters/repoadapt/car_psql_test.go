package repoadapt

import (
	"context"
	dbsql "database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
	_ "github.com/yawnak/fuel-record-crud/ent/runtime"
	"github.com/yawnak/fuel-record-crud/internal/services/vehicles"
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
	return ent.NewClient(ent.Driver(drv), ent.Debug(), ent.Log(log.Println))
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
	fmt.Print("car creation:")
	_, err = client.Car.Create().SetID(uuid.New()).SetMake("Shelby").SetModel("GT500").SetYear(2020).Save(ctx)
	if err != nil {
		t.Fatal("error creating car:", err)
	}

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

	firstFuelBuilder := client.FuelRecord.Create().SetID(fuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(10).SetDifference(10).SetCreatedAt(time.Now())
	secondFuelBuilder := client.FuelRecord.Create().SetID(nextFuelUUID).SetCarID(carUUID).SetCurrentFuelLiters(20).SetDifference(10).SetCreatedAt(time.Now()).SetPrevID(fuelUUID)

	fuelRecords, err := client.FuelRecord.CreateBulk(firstFuelBuilder, secondFuelBuilder).Save(ctx)
	if err != nil {
		t.Fatal("error creating fuel record bulk:", err)
	}

	_, err = client.OdometerRecord.Create().SetID(odometerUUID).SetCarID(carUUID).SetCurrentFuelLiters(100).SetDifference(100).SetCreatedAt(time.Now()).Save(ctx)
	if err != nil {
		t.Fatal("error creating odometer record:", err)
	}
	_, err = client.OdometerRecord.Create().SetID(nextOdometerUUID).SetCarID(carUUID).SetCurrentFuelLiters(200).SetDifference(100).SetCreatedAt(time.Now()).SetPrevID(odometerUUID).Save(ctx)
	if err != nil {
		t.Fatal("error creating next odometer record:", err)
	}

	fmt.Println("query next:")
	fmt.Println(client.FuelRecord.Query().Where(sql.FieldEQ(fuelrecord.FieldID, nextFuelUUID)).First(ctx))
	fmt.Println(client.FuelRecord.QueryPrev(fuelRecords[1]).All(ctx))

	//fmt.Println(client.Car.QueryFuelRecords(res).Order(fuelrecord.ByCreatedAt(entsql.OrderDesc())))
}
func Stuff() {
	client := &Client{}
	vehicles.NewVehicleService(client)
}
