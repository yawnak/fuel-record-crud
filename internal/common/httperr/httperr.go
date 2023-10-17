package httperr

import (
	"encoding/json"
	"net/http"
)

var (
	TypeInternal = "internal"
	TypeUnknown  = "unknown"
)

type SError struct {
	Code           int    `json:"code"`
	ErrType        string `json:"error_type"`
	DisplayMessage string `json:"message"`
	Err            error  `json:"-"`
}

func (e *SError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func NewInternal(err error) *SError {
	return &SError{
		Code:    http.StatusInternalServerError,
		ErrType: TypeInternal,
		Err:     err,
	}
}
