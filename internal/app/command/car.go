package command

import (
	"context"
	"time"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/vehicle"
)

type Vehicle struct {
	Create *CreateVehicleHandler
}

type VehicleModel interface {
	VehicleCreateService
}

func NewCarCommand(model VehicleModel) *Vehicle {
	return &Vehicle{
		Create: NewCreateVehicle(model),
	}
}

type CarCreateModel interface {
	CreateCar(context.Context, car.Car) (car.Car, error)
}

type VehicleCreateService interface {
	CreateVehicle(ctx context.Context,
		model, make string, year int32,
		initFuel *float64, fuelCreationTime time.Time,
		initOdometer *float64, odometerCreationTime time.Time,
	) (vehicle.Vehicle, error)
}

type CreateVehicle struct {
	Model string
	Make  string
	Year  int32
}

type CreateVehicleHandler struct {
	vehicleCreateService VehicleCreateService
}

func NewCreateVehicle(createService VehicleCreateService) *CreateVehicleHandler {
	return &CreateVehicleHandler{
		vehicleCreateService: createService,
	}
}

// func (c *CreateVehicleHandler) Handle(ctx context.Context, cmd CreateCar) (car.Car, error) {
// 	car.New(cmd.Model, cmd.Make, cmd.Year)
// 	return c.carCreateModel.CreateCar(ctx, car)
// }
