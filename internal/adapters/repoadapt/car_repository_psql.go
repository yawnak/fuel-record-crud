package repoadapt

import (
	"context"
	"fmt"

	dbsql "database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/internal/domain/car"
)

func EntCarToCar(entCar *ent.Car) car.Car {
	return car.UnmarshalCarFromDatabase(
		entCar.ID,
		entCar.Make,
		entCar.Model,
		entCar.Year,
	)
}

func CarToCreate(c car.Car, carCreator *ent.CarCreate) *ent.CarCreate {
	return carCreator.
		SetID(c.Id()).
		SetMake(c.Make()).
		SetModel(c.Model()).
		SetYear(c.Year())
}

func CarToUpdateOne(c car.Car, carUpdater *ent.CarUpdateOne) *ent.CarUpdateOne {
	return carUpdater.
		SetMake(c.Make()).
		SetModel(c.Model()).
		SetYear(c.Year())
}

type CarRepositoryPSQL struct {
	client *ent.CarClient
}

func NewCarRepositoryPSQL(user, password, host, port, dbname string) (*CarRepositoryPSQL, error) {
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	db, err := dbsql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv)).Car
	return &CarRepositoryPSQL{client: client}, nil
}

func (repo *CarRepositoryPSQL) CreateCar(ctx context.Context, c car.Car) (car.Car, error) {
	newCar, err := CarToCreate(c, repo.client.Create()).Save(ctx)
	if err != nil {
		return car.Car{}, err
	}
	return EntCarToCar(newCar), nil
}

func (repo *CarRepositoryPSQL) ReadCar(ctx context.Context, id uuid.UUID) (car.Car, error) {
	getCar, err := repo.client.Get(ctx, id)
	if err != nil {
		return car.Car{}, err
	}
	return EntCarToCar(getCar), nil
}

func (repo *CarRepositoryPSQL) UpdateCar(ctx context.Context, c car.Car) (car.Car, error) {
	updatedCar, err := CarToUpdateOne(c, repo.client.UpdateOneID(c.Id())).Save(ctx)
	if err != nil {
		return car.Car{}, err
	}
	return EntCarToCar(updatedCar), nil
}

func (repo *CarRepositoryPSQL) DeleteCar(ctx context.Context, id uuid.UUID) error {
	return repo.client.DeleteOneID(id).Exec(ctx)
}
