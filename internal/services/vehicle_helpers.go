package services

import (
	"github.com/yawnak/fuel-record-crud/internal/domain/car"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
	"github.com/yawnak/fuel-record-crud/internal/domain/vehicle"
	"github.com/yawnak/fuel-record-crud/pkg/history"
)

func VehicleFromOneRecord(cr car.Car, fuelRec *record.FuelGauge, odometerRec *record.Odometer) (vehicle.Vehicle, error) {
	return vehicle.BuildVehicle(
		cr,
		record.FuelGaugeHistory{History: history.NewHistory(fuelRec)},
		record.OdometerHistory{History: history.NewHistory(odometerRec)}), nil
}
