package hooks

import (
	"entgo.io/ent"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
)

type fuelRecordHooks struct {
	setNextFields []string
}

var (
	FuelRecord = fuelRecordHooks{
		setNextFields: []string{
			fuelrecord.EdgeNext, fuelrecord.FieldNextFuelRecordID,
		},
	}
)

func (hook fuelRecordHooks) ForbidSetNext(err error) ent.Hook {
	return HasFields(err,
		hook.setNextFields...)
}
