package rules

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
)

type fuelRecord struct {
	setNextFields []string
}

var (
	FuelRecord = fuelRecord{
		setNextFields: []string{
			fuelrecord.EdgeNext, fuelrecord.FieldNextFuelRecordID,
		},
	}
)

func (fuelRecord) DenySetNext(err error) privacy.MutationRuleFunc {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(ent.OpUpdateOne) && HasFields(m, err, FuelRecord.setNextFields...) {
			return privacy.Denyf("", err)
		}
		return privacy.Skip
	})
}
