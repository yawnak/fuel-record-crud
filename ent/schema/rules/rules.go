package rules

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

func DenyDelete(err error) privacy.MutationRuleFunc {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(ent.OpDelete) {
			return privacy.Denyf("", err)
		}
		return privacy.Skip
	})
}
