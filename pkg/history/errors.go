package history

import "errors"

var (
	ErrHistoryEmpty  = errors.New("history is empty")
	ErrHeadIsNotLast = errors.New("head is not last")
)
