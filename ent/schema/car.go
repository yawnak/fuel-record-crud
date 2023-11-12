package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("car_id").Immutable().Default(uuid.New),
		field.String("make").NotEmpty(),
		field.String("model").NotEmpty(),
		field.Int32("year"),
	}
}

func (Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fuel_records", FuelRecord.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("odometer_records", OdometerRecord.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
