package repoadapt

import (
	"context"

	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/internal/domain/event"
)

type OdometerRecordRepoPSQL struct {
	client *ent.OdometerRecordClient
}

func (repo *OdometerRecordRepoPSQL) Create(ctx context.Context, odometerRecord event.OdometerIncrease) (event.OdometerIncrease, error) {
	return event.OdometerIncrease{}, nil
}
