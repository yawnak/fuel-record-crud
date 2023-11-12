package command

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/vehicle"
)

type Vehicle struct {
	Create *CreateVehicleHandler
}

type VehicleModel interface {
	VehicleCreateService
}

func NewVehicleCommand(model VehicleModel) *Vehicle {
	return &Vehicle{
		Create: NewCreateVehicle(model),
	}
}

type VehicleCreateService interface {
	CreateVehicle(ctx context.Context,
		model, make string, year int32,
		initFuel *float64, fuelCreationTime time.Time,
		initOdometer *float64, odometerCreationTime time.Time,
	) (*vehicle.Vehicle, error)
}

type CreateVehicle struct {
	Model                string
	Make                 string
	Year                 int32
	CurrentFuel          *float64
	FuelCreationTime     *time.Time
	CurrentOdometer      *float64
	OdometerCreationTime *time.Time
}

type CreateVehicleHandler struct {
	vehicleCreateService VehicleCreateService
}

func NewCreateVehicle(createService VehicleCreateService) *CreateVehicleHandler {
	return &CreateVehicleHandler{
		vehicleCreateService: createService,
	}
}

func (c *CreateVehicleHandler) Handle(ctx context.Context, cmd CreateVehicle) (*vehicle.Vehicle, error) {
	return c.vehicleCreateService.CreateVehicle(ctx, cmd.Model, cmd.Make, cmd.Year,
		cmd.CurrentFuel, *cmd.FuelCreationTime,
		cmd.CurrentOdometer, *cmd.OdometerCreationTime)
}

type CreateFuelRecord struct {
	VehicleID          uuid.UUID
	CurrentFuel        float64
	FuelRecordDatetime time.Time
}

type CreateOdometerRecord struct {
	VehicleID              uuid.UUID
	CurrentOdometer        float64
	OdometerRecordDatetime time.Time
}
