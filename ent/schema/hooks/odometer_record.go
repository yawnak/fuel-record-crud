package hooks

import (
	"entgo.io/ent"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
)

type odometerRecordHooks struct {
	setNextFields []string
}

var (
	OdometerRecord = odometerRecordHooks{
		setNextFields: []string{
			odometerrecord.EdgeNext, odometerrecord.FieldNextOdometerRecordID,
		},
	}
)

var ()

func (hook odometerRecordHooks) ForbidSetNext(err error) ent.Hook {
	return HasFields(err,
		hook.setNextFields...)
}
