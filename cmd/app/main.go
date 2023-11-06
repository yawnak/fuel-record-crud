package main

import (
	"log"

	_ "github.com/yawnak/fuel-record-crud/ent/runtime"
	"github.com/yawnak/fuel-record-crud/internal/app"
	"github.com/yawnak/fuel-record-crud/internal/app/command"
	"github.com/yawnak/fuel-record-crud/internal/app/query"
	"github.com/yawnak/fuel-record-crud/internal/common/server"
	"github.com/yawnak/fuel-record-crud/internal/config"
	"github.com/yawnak/fuel-record-crud/internal/routes"
	"github.com/yawnak/fuel-record-crud/internal/services"
	"github.com/yawnak/fuel-record-crud/web"
)

var (
	configPath string = "configs/development.local.yaml"
)

func main() {
	repo := initRepo()
	serverConf := Must(config.NewServerConfig(configPath))

	vehicleService := services.NewVehicleService(repo)

	app := app.Application{
		Commands: app.Commands{
			Vehicle: command.NewVehicleCommand(vehicleService),
		},
		Queries: app.Queries{
			Vehicle: query.NewVehicleQuery(vehicleService),
		},
	}

	httpWrapper := routes.NewHTTPWrapper(&app)

	err := server.RunHTTPServer(
		serverConf.Port,
		web.Templates,
		routes.InitRoutes(httpWrapper))

	log.Panicln("error running server:", err)
}

func Must[T any](t T, err error) T {
	if err != nil {
		log.Panicln(err)
	}
	return t
}
