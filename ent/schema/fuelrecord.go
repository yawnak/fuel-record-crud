package schema

import "entgo.io/ent"

// FuelRecord holds the schema definition for the FuelRecord entity.
type FuelRecord struct {
	ent.Schema
}

// Fields of the FuelRecord.
func (FuelRecord) Fields() []ent.Field {
	return nil
}

// Edges of the FuelRecord.
func (FuelRecord) Edges() []ent.Edge {
	return nil
}
