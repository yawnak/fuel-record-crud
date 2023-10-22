package app

import (
	"github.com/yawnak/fuel-record-crud/internal/app/command"
	"github.com/yawnak/fuel-record-crud/internal/app/query"
)

type Application struct {
	Commands
	Queries
}

type Commands struct {
	Car command.Vehicle
}

type Queries struct {
	Car query.Car
}
