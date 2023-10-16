package httperr

import "net/http"

var (
	TypeInternal = "internal"
	TypeUnknown  = "unknown"
)

type SError struct {
	Code           int
	ErrType        string
	DisplayMessage string
	Err            error
}

func (e SError) Error() string {
	return e.Err.Error()
}

func NewInternal(err error) SError {
	return SError{
		Code:    http.StatusInternalServerError,
		ErrType: TypeInternal,
		Err:     err,
	}
}
