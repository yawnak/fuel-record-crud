package hooks

import (
	"context"

	"entgo.io/ent"
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/ent/hook"
)

// checks fields, if they are set, returns error
func HasFields(err error, fields ...string) ent.Hook {
	return hook.If(hook.FixedError(err),
		func(ctx context.Context, m ent.Mutation) bool {
			_, ok := lo.Find(fields, func(field string) bool {
				_, ok := m.Field(field)
				return ok
			})
			return ok
		},
	)
}
