package history

import (
	"fmt"

	"github.com/samber/lo"
)

type Historable[T any] interface {
	comparable
	NewNext() (T, error)
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

func (history *History[H]) AddRecord() error {
	head, err := lo.Last(history.list)
	if err != nil || lo.IsEmpty(head) {
		return ErrHistoryEmpty
	}
	if !head.IsLast() {
		return ErrHeadIsNotLast
	}
	newHead, err := head.NewNext()
	if err != nil {
		return fmt.Errorf("error in head.NewNext: %w", err)
	}

	history.list = append(history.list, newHead)

	return nil
}

func (history *History[H]) IsComplete() bool {
	return len(history.list) > 0 &&
		history.list[0].IsFirst() &&
		history.list[len(history.list)-1].IsLast()
}
