package command

import (
	"context"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
)

type Vehicle struct {
	Create *CreateVehicleHandler
}

type VehicleModel interface {
	CarCreateModel
}

func NewCarCommand(model VehicleModel) *Vehicle {
	return &Vehicle{
		Create: NewCreateCar(model),
	}
}

type CarCreateModel interface {
	CreateCar(context.Context, car.Car) (car.Car, error)
}

type CreateVehicle struct {
	Model string
	Make  string
	Year  int32
}

type CreateVehicleHandler struct {
	carCreateModel CarCreateModel
}

func NewCreateCar(createModel CarCreateModel) *CreateVehicleHandler {
	return &CreateVehicleHandler{
		carCreateModel: createModel,
	}
}

func (c *CreateVehicleHandler) Handle(ctx context.Context, cmd CreateCar) (car.Car, error) {
	car.New(cmd.Model, cmd.Make, cmd.Year)
	return c.carCreateModel.CreateCar(ctx, car)
}
