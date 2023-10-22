package erradapt

import (
	"errors"
	"net/http"

	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
	"github.com/yawnak/fuel-record-crud/internal/domain/view"
)

var (
	TypeCreateRecordConflict = "create_record_conflict"
)

func adaptErrNextRecordAlreadyExists(err error) *httperr.SError {
	if !errors.Is(err, view.ErrFuelGaugeRecordNotLast) {
		return nil
	}
	return &httperr.SError{
		Code:           http.StatusConflict,
		ErrType:        TypeCreateRecordConflict,
		DisplayMessage: "next record already exists",
		Err:            view.ErrFuelGaugeRecordNotLast,
	}
}

func registerCarAdapters(adapter *Adapter) {
	adapter.registerAdapterFunc(adaptErrNextRecordAlreadyExists)
}
