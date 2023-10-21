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
		field.UUID("car_id", uuid.UUID{}).Immutable(),
		field.UUID("next_fuel_record_id", uuid.UUID{}).Optional().Immutable(),
	}
}

// Edges of the FuelRecord.
func (FuelRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("fuel_records").
			Unique().Required().Immutable().Field("car_id"),
		edge.To("prev", FuelRecord.Type).Unique().Immutable().
			From("next").Unique().Immutable().Field("next_fuel_record_id"),
	}
}
