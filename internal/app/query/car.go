package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
)

type Car struct {
	GetHandler *GetСarHandler
}

type CarModel interface {
	CarReadModel
}

func NewCarQuery(model CarModel) *Car {
	return &Car{
		GetHandler: NewCarHandler(model),
	}
}

type CarReadModel interface {
	GetCar(ctx context.Context, id uuid.UUID) (car.Car, error)
}

type GetСarHandler struct {
	carReadModel CarReadModel
}

func NewCarHandler(readModel CarReadModel) *GetСarHandler {
	return &GetСarHandler{
		carReadModel: readModel,
	}
}

func (h *GetСarHandler) Handle(ctx context.Context, id uuid.UUID) (car.Car, error) {
	return h.carReadModel.GetCar(ctx, id)
}
