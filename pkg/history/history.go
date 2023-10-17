package history

import (
	"fmt"

	"github.com/samber/lo"
)

type Historable[T any] interface {
	comparable
	NewNext(difference float64) (T, error)
	IsFirst() bool
	IsLast() bool
}

type History[H Historable[H]] struct {
	list []H
}

func NewHistory[H Historable[H]](record H) *History[H] {
	return &History[H]{
		list: []H{record},
	}
}

func (history *History[H]) AddRecord(difference float64) error {
	head, err := lo.Last(history.list)
	if err != nil || lo.IsEmpty(head) {
		return ErrHistoryEmpty
	}
	if !head.IsLast() {
		return ErrHeadIsNotLast
	}
	newHead, err := head.NewNext(difference)
	if err != nil {
		return fmt.Errorf("error in head.NewNext: %w", err)
	}

	history.list = append(history.list, newHead)

	return nil
}
