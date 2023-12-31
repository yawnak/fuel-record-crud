package erradapt

import (
	"errors"
	"net/http"

	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
	"github.com/yawnak/fuel-record-crud/internal/domain/record"
)

var (
	TypeCreateRecordConflict = "create_record_conflict"
)

func adaptErrNextRecordAlreadyExists(err error) *httperr.SError {
	if !errors.Is(err, record.ErrRecordNotLast) {
		return nil
	}
	return &httperr.SError{
		Code:           http.StatusConflict,
		ErrType:        TypeCreateRecordConflict,
		DisplayMessage: "next record already exists",
		Err:            err,
	}
}

func registerCarAdapters(adapter *Adapter) {
	adapter.registerAdapterFunc(adaptErrNextRecordAlreadyExists)
}
