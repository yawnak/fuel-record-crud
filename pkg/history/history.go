package history

import (
	"fmt"

	"github.com/samber/lo"
)

type Historable[T any, Event any] interface {
	comparable
	Copier[T]
	NewNext(e Event) (T, error)
	IsFirst() bool
	IsLast() bool
}

type Copier[T any] interface {
	Copy() T
}

type History[H Historable[H, Event], Event any] struct {
	list []H
}

func NewHistory[H Historable[H, Event], Event any](record H) History[H, Event] {
	if lo.IsEmpty(record) {
		return History[H, Event]{
			list: []H{},
		}
	}
	return History[H, Event]{
		list: []H{record},
	}
}

func UnmarshalHistory[H Historable[H, Event], Event any](list []H) History[H, Event] {
	return History[H, Event]{
		list: list,
	}
}

func (history *History[H, E]) AddRecord(event E) error {
	head, err := lo.Last(history.list)
	if err != nil || lo.IsEmpty(head) {
		return ErrHistoryEmpty
	}
	if !head.IsLast() {
		return ErrHeadIsNotLast
	}
	newHead, err := head.NewNext(event)
	if err != nil {
		return fmt.Errorf("error in head.NewNext: %w", err)
	}

	history.list = append(history.list, newHead)

	return nil
}

func (history *History[H, E]) IsComplete() bool {
	return len(history.list) > 0 &&
		history.list[0].IsFirst() &&
		history.list[len(history.list)-1].IsLast()
}

// returns a copy of the head record
func (history *History[H, E]) Head() H {
	if len(history.list) == 0 {
		return *new(H)
	}
	return history.list[len(history.list)-1].Copy()
}
