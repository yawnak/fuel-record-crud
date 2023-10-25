package domain

import (
	"context"

	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
)

type Beginner interface {
	BeginTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	Commit() error
	Rollback() error
}

type TxWithRepos interface {
	Tx
	CarRepo() car.CarRepo
	FuelGaugeChangeRepo() record.FuelGaugeRepo
	OdometerIncreaseRepo() record.OdometerRepo
}
