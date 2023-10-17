package car

import "errors"

// record errors
var (
	ErrNextRecordAlreadyExists = errors.New("next record already exists")
)

// odometer errors
var (
	ErrOdometerNegativeDifference = errors.New("odometer negative difference")
)

var (

)
