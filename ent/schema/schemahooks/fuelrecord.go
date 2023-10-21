package schemahooks

import (
	"context"

	"entgo.io/ent"
	gen "github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/ent/hook"
)

func EnsureIsLast() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.FuelRecordFunc(func(ctx context.Context, m *gen.FuelRecordMutation) (ent.Value, error) {
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpDeleteOne)
}
