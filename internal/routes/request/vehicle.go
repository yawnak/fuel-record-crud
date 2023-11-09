package request

import (
	"time"

	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/internal/app/command"
)

type DateTime time.Time

func (ct *DateTime) UnmarshalParam(param string) error {
	if param == "" {
		return nil
	}
	t, err := time.Parse(time.DateOnly, param)
	if err != nil {
		return err
	}
	*ct = DateTime(t)
	return nil
}

type CreateVehicle struct {
	Make                 string    `form:"make"`
	Model                string    `form:"model"`
	Year                 int32     `form:"year"`
	CurrentFuel          *float64  `form:"current_fuel"`
	FuelCreationTime     *DateTime `form:"fuel_record_date"`
	CurrentOdometer      *float64  `form:"current_odometer"`
	OdometerCreationTime *DateTime `form:"odometer_record_date"`
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
