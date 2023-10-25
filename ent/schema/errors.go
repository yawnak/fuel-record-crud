package schema

import "errors"

var (
	ErrSetNextIsForbidden = errors.New("set next is forbidden")
	ErrDeleteIsForbidden  = errors.New("delete is forbidden")
)
