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
		field.UUID("next_fuel_record_id", uuid.NullUUID{}).Optional(),
	}
}

// Edges of the FuelRecord.
func (FuelRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("fuel_records").Unique().Required(),
		edge.To("next", FuelRecord.Type).Unique().Immutable().
			From("next").Unique().Immutable(),
	}
}
