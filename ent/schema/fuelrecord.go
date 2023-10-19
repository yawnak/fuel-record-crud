package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FuelRecord holds the schema definition for the FuelRecord entity.
type FuelRecord struct {
	ent.Schema
}

// Fields of the FuelRecord.
func (FuelRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("fuel_record_id").Immutable(),
		field.Float("current_fuel_liters").Positive(),
		field.Float("difference"),
		field.Time("created_at").Immutable(),
	}
}

// Edges of the FuelRecord.
func (FuelRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("fuel_records").Unique().Required(),
		edge.To("next", FuelRecord.Type).Unique().From("prev").Unique(),
	}
}
