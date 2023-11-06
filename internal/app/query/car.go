package query

import (
	"context"

	"github.com/yawnak/fuel-record-crud/internal/domain/vehicle"
)

type Vehicle struct {
	GetVehiclesCurrentData *GetVehiclesCurrentDataHandler
}

type VehicleModel interface {
	VehicleCurrentReadModel
}

func NewVehicleQuery(model VehicleModel) *Vehicle {
	return &Vehicle{
		GetVehiclesCurrentData: NewGetVehiclesCurrentDataHandler(model),
	}
}

type VehicleCurrentReadModel interface {
	GetVehiclesCurrent(ctx context.Context) ([]vehicle.Vehicle, error)
}

type GetVehiclesCurrentDataHandler struct {
	model VehicleCurrentReadModel
}

func NewGetVehiclesCurrentDataHandler(readModel VehicleCurrentReadModel) *GetVehiclesCurrentDataHandler {
	return &GetVehiclesCurrentDataHandler{
		model: readModel,
	}
}

func (h *GetVehiclesCurrentDataHandler) Handle(ctx context.Context) ([]vehicle.Vehicle, error) {
	return h.model.GetVehiclesCurrent(ctx)
}
