package erradapt

import (
	"net/http"

	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
)

var (
	TypeCreateRecordConflict = "create_record_conflict"
)

func adaptErrNextRecordAlreadyExists(err error) *httperr.SError {
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
