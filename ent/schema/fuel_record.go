package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/privacy"
	"github.com/yawnak/fuel-record-crud/ent/schema/rules"
)

// FuelRecord holds the schema definition for the FuelRecord entity.
type FuelRecord struct {
	ent.Schema
}

// Fields of the FuelRecord.
func (FuelRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("fuel_record_id").Immutable(),
		field.Float("current_fuel_liters").Min(0).Immutable(),
		field.Float("difference").Immutable(),
		field.Time("created_at").Immutable(),
		field.UUID("car_id", uuid.UUID{}).Immutable(),
		field.UUID("next_fuel_record_id", uuid.UUID{}).Optional().Nillable().Immutable(),
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

// func (FuelRecord) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hooks.FuelRecord.ForbidSetNext(ErrSetNextIsForbidden),
// 	}
// }

func (FuelRecord) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.DenyDelete(ErrDeleteIsForbidden),
			rules.FuelRecord.DenySetNext(ErrSetNextIsForbidden),
		},
	}
}
