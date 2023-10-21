package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OdometerRecord holds the schema definition for the OdometerRecord entity.
type OdometerRecord struct {
	ent.Schema
}

// Fields of the OdometerRecord.
func (OdometerRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("odometer_record_id").Immutable(),
		field.Float("current_fuel_liters").Positive(),
		field.Float("difference").Positive(),
		field.Time("created_at").Immutable(),
	}
}

// Edges of the OdometerRecord.
func (OdometerRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("odometer_records").Unique().Required(),
		edge.To("prev", OdometerRecord.Type).Unique().From("next").Unique().Immutable(),
	}
}
