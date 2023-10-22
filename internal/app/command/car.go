package command

import (
	"context"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
)

type Car struct {
	Create *CreateCarHandler
}

type CarModel interface {
	CarCreateModel
}

func NewCarCommand(model CarModel) *Car {
	return &Car{
		Create: NewCreateCar(model),
	}
}

type CarCreateModel interface {
	CreateCar(context.Context, car.Car) (car.Car, error)
}

type CreateCarHandler struct {
	carCreateModel CarCreateModel
}

func NewCreateCar(createModel CarCreateModel) *CreateCarHandler {
	return &CreateCarHandler{
		carCreateModel: createModel,
	}
}

func (c *CreateCarHandler) Handle(ctx context.Context, car car.Car) (car.Car, error) {
	return c.carCreateModel.CreateCar(ctx, car)
}
