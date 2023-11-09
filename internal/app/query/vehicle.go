package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/internal/domain/vehicle"
)

type Vehicle struct {
	GetVehiclesCurrentData *GetVehiclesCurrentDataHandler
	GetVehicle             *GetVehicleHandler
}

type VehicleModel interface {
	VehiclesCurrentReadModel
	VehicleReadModel
}

func NewVehicleQuery(model VehicleModel) *Vehicle {
	return &Vehicle{
		GetVehiclesCurrentData: NewGetVehiclesCurrentDataHandler(model),
		GetVehicle:             NewGetVehicleHandler(model),
	}
}

// ===== GET VEHICLES CURRENT DATA =====

type VehiclesCurrentReadModel interface {
	GetVehiclesCurrent(ctx context.Context) ([]vehicle.Vehicle, error)
}

type GetVehiclesCurrentDataHandler struct {
	model VehiclesCurrentReadModel
}

func NewGetVehiclesCurrentDataHandler(readModel VehiclesCurrentReadModel) *GetVehiclesCurrentDataHandler {
	return &GetVehiclesCurrentDataHandler{
		model: readModel,
	}
}

func (h *GetVehiclesCurrentDataHandler) Handle(ctx context.Context) ([]vehicle.Vehicle, error) {
	return h.model.GetVehiclesCurrent(ctx)
}

// ===== GET VEHICLE =====

type VehicleReadModel interface {
	GetVehicle(ctx context.Context, id uuid.UUID) (*vehicle.Vehicle, error)
}

type GetVehicleHandler struct {
	model VehicleReadModel
}

func NewGetVehicleHandler(readModel VehicleReadModel) *GetVehicleHandler {
	return &GetVehicleHandler{
		model: readModel,
	}
}

func (h *GetVehicleHandler) Handle(ctx context.Context, vehicleId uuid.UUID) (*vehicle.Vehicle, error) {
	return h.model.GetVehicle(ctx, vehicleId)
}
