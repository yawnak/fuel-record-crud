package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/schema/hooks"
)

// OdometerRecord holds the schema definition for the OdometerRecord entity.
type OdometerRecord struct {
	ent.Schema
}

// Fields of the OdometerRecord.
func (OdometerRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("odometer_record_id").Immutable(),
		field.Float("current_fuel_liters").Min(0).Immutable(),
		field.Float("difference").Min(0).Immutable(),
		field.Time("created_at").Immutable(),
		field.UUID("car_id", uuid.UUID{}).Immutable(),
		field.UUID("next_odometer_record_id", uuid.UUID{}).Optional().Immutable(),
	}
}

// Edges of the OdometerRecord.
func (OdometerRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("odometer_records").
			Unique().Required().Field("car_id").Immutable(),
		edge.To("prev", OdometerRecord.Type).Unique().Immutable().
			From("next").Unique().Immutable().Field("next_odometer_record_id"),
	}
}

func (OdometerRecord) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.OdometerRecord.ForbidSetNext(ErrSetNextIsForbidden),
	}
}
