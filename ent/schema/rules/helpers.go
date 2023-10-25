package rules

import (
	"entgo.io/ent"
	"github.com/samber/lo"
)

// checks fields, if they are set, returns error
func HasFields(m ent.Mutation, err error, fields ...string) bool {
	_, ok := lo.Find(fields, func(field string) bool {
		_, ok := m.Field(field)
		return ok
	})
	return ok
}
