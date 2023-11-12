package repoadapt

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/ent"
	entcar "github.com/yawnak/fuel-record-crud/ent/car"
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

func CarToCreate(c *car.Car, carCreator *ent.CarCreate) *ent.CarCreate {
	return carCreator.
		SetID(c.Id()).
		SetMake(c.Make()).
		SetModel(c.Model()).
		SetYear(c.Year())
}

func CarToUpdateOne(c *car.Car, carUpdater *ent.CarUpdateOne) *ent.CarUpdateOne {
	return carUpdater.
		SetMake(c.Make()).
		SetModel(c.Model()).
		SetYear(c.Year())
}

type CarRepositoryPSQL struct {
	client *ent.CarClient
}

func NewCarRepositoryPSQL(carClient *ent.CarClient) *CarRepositoryPSQL {
	return &CarRepositoryPSQL{
		client: carClient,
	}
}

func (repo *CarRepositoryPSQL) CreateCar(ctx context.Context, c *car.Car) error {
	_, err := CarToCreate(c, repo.client.Create()).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CarRepositoryPSQL) GetCar(ctx context.Context, id uuid.UUID) (car.Car, error) {
	getCar, err := repo.client.Get(ctx, id)
	if err != nil {
		return car.Car{}, err
	}
	return EntCarToCar(getCar), nil
}

func (repo *CarRepositoryPSQL) QueryCars(ctx context.Context) ([]car.Car, error) {
	getCars, err := repo.client.Query().Order(entcar.ByCreateTime(sql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(getCars, func(getCar *ent.Car, _ int) car.Car {
		return EntCarToCar(getCar)
	}), nil
}

func (repo *CarRepositoryPSQL) UpdateCar(ctx context.Context, c *car.Car) (car.Car, error) {
	updatedCar, err := CarToUpdateOne(c, repo.client.UpdateOneID(c.Id())).Save(ctx)
	if err != nil {
		return car.Car{}, err
	}
	return EntCarToCar(updatedCar), nil
}

func (repo *CarRepositoryPSQL) DeleteCar(ctx context.Context, id uuid.UUID) error {
	return repo.client.DeleteOneID(id).Exec(ctx)
}
