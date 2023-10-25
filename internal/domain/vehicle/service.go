package vehicle

type VehicleService interface {
	Create(model, make string, year int32, initFuel *float64, initOdometer *float64) error
}
