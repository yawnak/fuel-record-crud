package record

import "errors"

var (
	ErrRecordStatesDontMatch    = errors.New("record states don't match")
	ErrNewNextIsEarlierThenLast = errors.New("new next is earlier then last")
)
