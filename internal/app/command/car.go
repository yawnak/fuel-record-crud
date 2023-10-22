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

type CreateVehicleHandler struct {
	carCreateModel CarCreateModel
}

func NewCreateCar(createModel CarCreateModel) *CreateVehicleHandler {
	return &CreateVehicleHandler{
		carCreateModel: createModel,
	}
}

func (c *CreateVehicleHandler) Handle(ctx context.Context, car car.Car) (car.Car, error) {
	return c.carCreateModel.CreateCar(ctx, car)
}
