package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("car_id").Immutable(),
		field.String("make").NotEmpty(),
		field.String("model").NotEmpty(),
		field.Int8("year").Positive(),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fuel_records", FuelRecord.Type),
		edge.To("odometer_records", OdometerRecord.Type),
	}
}