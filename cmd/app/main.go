package main

import (
	"fmt"

	"github.com/yawnak/fuel-record-crud/internal/adapters/repoadapt"
	"github.com/yawnak/fuel-record-crud/internal/app/command"
	"github.com/yawnak/fuel-record-crud/internal/services/vehicles"
)

func main() {
	fmt.Println("Hello world!")
	client := &repoadapt.Client{}
	serv := vehicles.NewVehicleService(client)
	command.NewCarCommand(serv)
}
