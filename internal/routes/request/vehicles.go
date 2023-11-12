package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/internal/app/command"
)

type CreateVehicle struct {
	Make                 string    `form:"make"`
	Model                string    `form:"model"`
	Year                 int32     `form:"year"`
	CurrentFuel          *float64  `form:"current_fuel"`
	FuelCreationTime     *DateTime `form:"fuel_record_date"` // 2006-01-02T15:04
	CurrentOdometer      *float64  `form:"current_odometer"`
	OdometerCreationTime *DateTime `form:"odometer_record_date"` // 2006-01-02T15:04
}

func (req *CreateVehicle) ToCmd() command.CreateVehicle {
	return command.CreateVehicle{
		Make:                 req.Make,
		Model:                req.Model,
		Year:                 req.Year,
		CurrentFuel:          req.CurrentFuel,
		FuelCreationTime:     lo.ToPtr(time.Time(lo.FromPtr(req.FuelCreationTime))),
		CurrentOdometer:      req.CurrentOdometer,
		OdometerCreationTime: lo.ToPtr(time.Time(lo.FromPtr(req.OdometerCreationTime))),
	}
}

type CreateFuelRecord struct {
	VehicleID    string   `form:"vehicle_id"`
	CurrentFuel  float64  `form:"current_fuel"`
	FuelDatetime DateTime `form:"fuel_record_date"` // 2006-01-02T15:04
}

func (req *CreateFuelRecord) ToCmd() (command.CreateFuelRecord, error) {
	vehid, err := uuid.Parse(req.VehicleID)
	if err != nil {
		return command.CreateFuelRecord{}, err
	}
	return command.CreateFuelRecord{
		VehicleID:          vehid,
		CurrentFuel:        req.CurrentFuel,
		FuelRecordDatetime: time.Time(req.FuelDatetime),
	}, nil
}

type CreateOdometerRecord struct {
	VehicleID        string   `form:"vehicle_id"`
	CurrentOdometer  float64  `form:"current_odometer"`
	OdometerDatetime DateTime `form:"odometer_record_date"` // 2006-01-02T15:04
}

func (req *CreateOdometerRecord) ToCmd() (command.CreateOdometerRecord, error) {
	vehid, err := uuid.Parse(req.VehicleID)
	if err != nil {
		return command.CreateOdometerRecord{}, err
	}
	return command.CreateOdometerRecord{
		VehicleID:              vehid,
		CurrentOdometer:        req.CurrentOdometer,
		OdometerRecordDatetime: time.Time(req.OdometerDatetime),
	}, nil
}
