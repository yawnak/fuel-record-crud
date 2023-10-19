package schema

import "entgo.io/ent"

// OdometerRecord holds the schema definition for the OdometerRecord entity.
type OdometerRecord struct {
	ent.Schema
}

// Fields of the OdometerRecord.
func (OdometerRecord) Fields() []ent.Field {
	return nil
}

// Edges of the OdometerRecord.
func (OdometerRecord) Edges() []ent.Edge {
	return nil
}
