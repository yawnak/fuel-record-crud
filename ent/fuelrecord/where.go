// Code generated by ent, DO NOT EDIT.

package fuelrecord

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLTE(FieldID, id))
}

// CurrentFuelLiters applies equality check predicate on the "current_fuel_liters" field. It's identical to CurrentFuelLitersEQ.
func CurrentFuelLiters(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldCurrentFuelLiters, v))
}

// Difference applies equality check predicate on the "difference" field. It's identical to DifferenceEQ.
func Difference(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldDifference, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldCreatedAt, v))
}

// CarID applies equality check predicate on the "car_id" field. It's identical to CarIDEQ.
func CarID(v uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldCarID, v))
}

// NextFuelRecordID applies equality check predicate on the "next_fuel_record_id" field. It's identical to NextFuelRecordIDEQ.
func NextFuelRecordID(v uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldNextFuelRecordID, v))
}

// CurrentFuelLitersEQ applies the EQ predicate on the "current_fuel_liters" field.
func CurrentFuelLitersEQ(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldCurrentFuelLiters, v))
}

// CurrentFuelLitersNEQ applies the NEQ predicate on the "current_fuel_liters" field.
func CurrentFuelLitersNEQ(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNEQ(FieldCurrentFuelLiters, v))
}

// CurrentFuelLitersIn applies the In predicate on the "current_fuel_liters" field.
func CurrentFuelLitersIn(vs ...float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIn(FieldCurrentFuelLiters, vs...))
}

// CurrentFuelLitersNotIn applies the NotIn predicate on the "current_fuel_liters" field.
func CurrentFuelLitersNotIn(vs ...float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotIn(FieldCurrentFuelLiters, vs...))
}

// CurrentFuelLitersGT applies the GT predicate on the "current_fuel_liters" field.
func CurrentFuelLitersGT(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGT(FieldCurrentFuelLiters, v))
}

// CurrentFuelLitersGTE applies the GTE predicate on the "current_fuel_liters" field.
func CurrentFuelLitersGTE(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGTE(FieldCurrentFuelLiters, v))
}

// CurrentFuelLitersLT applies the LT predicate on the "current_fuel_liters" field.
func CurrentFuelLitersLT(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLT(FieldCurrentFuelLiters, v))
}

// CurrentFuelLitersLTE applies the LTE predicate on the "current_fuel_liters" field.
func CurrentFuelLitersLTE(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLTE(FieldCurrentFuelLiters, v))
}

// DifferenceEQ applies the EQ predicate on the "difference" field.
func DifferenceEQ(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldDifference, v))
}

// DifferenceNEQ applies the NEQ predicate on the "difference" field.
func DifferenceNEQ(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNEQ(FieldDifference, v))
}

// DifferenceIn applies the In predicate on the "difference" field.
func DifferenceIn(vs ...float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIn(FieldDifference, vs...))
}

// DifferenceNotIn applies the NotIn predicate on the "difference" field.
func DifferenceNotIn(vs ...float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotIn(FieldDifference, vs...))
}

// DifferenceGT applies the GT predicate on the "difference" field.
func DifferenceGT(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGT(FieldDifference, v))
}

// DifferenceGTE applies the GTE predicate on the "difference" field.
func DifferenceGTE(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGTE(FieldDifference, v))
}

// DifferenceLT applies the LT predicate on the "difference" field.
func DifferenceLT(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLT(FieldDifference, v))
}

// DifferenceLTE applies the LTE predicate on the "difference" field.
func DifferenceLTE(v float64) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLTE(FieldDifference, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldLTE(FieldCreatedAt, v))
}

// CarIDEQ applies the EQ predicate on the "car_id" field.
func CarIDEQ(v uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldCarID, v))
}

// CarIDNEQ applies the NEQ predicate on the "car_id" field.
func CarIDNEQ(v uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNEQ(FieldCarID, v))
}

// CarIDIn applies the In predicate on the "car_id" field.
func CarIDIn(vs ...uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIn(FieldCarID, vs...))
}

// CarIDNotIn applies the NotIn predicate on the "car_id" field.
func CarIDNotIn(vs ...uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotIn(FieldCarID, vs...))
}

// NextFuelRecordIDEQ applies the EQ predicate on the "next_fuel_record_id" field.
func NextFuelRecordIDEQ(v uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldEQ(FieldNextFuelRecordID, v))
}

// NextFuelRecordIDNEQ applies the NEQ predicate on the "next_fuel_record_id" field.
func NextFuelRecordIDNEQ(v uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNEQ(FieldNextFuelRecordID, v))
}

// NextFuelRecordIDIn applies the In predicate on the "next_fuel_record_id" field.
func NextFuelRecordIDIn(vs ...uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIn(FieldNextFuelRecordID, vs...))
}

// NextFuelRecordIDNotIn applies the NotIn predicate on the "next_fuel_record_id" field.
func NextFuelRecordIDNotIn(vs ...uuid.UUID) predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotIn(FieldNextFuelRecordID, vs...))
}

// NextFuelRecordIDIsNil applies the IsNil predicate on the "next_fuel_record_id" field.
func NextFuelRecordIDIsNil() predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldIsNull(FieldNextFuelRecordID))
}

// NextFuelRecordIDNotNil applies the NotNil predicate on the "next_fuel_record_id" field.
func NextFuelRecordIDNotNil() predicate.FuelRecord {
	return predicate.FuelRecord(sql.FieldNotNull(FieldNextFuelRecordID))
}

// HasCar applies the HasEdge predicate on the "car" edge.
func HasCar() predicate.FuelRecord {
	return predicate.FuelRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CarTable, CarColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCarWith applies the HasEdge predicate on the "car" edge with a given conditions (other predicates).
func HasCarWith(preds ...predicate.Car) predicate.FuelRecord {
	return predicate.FuelRecord(func(s *sql.Selector) {
		step := newCarStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() predicate.FuelRecord {
	return predicate.FuelRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, NextTable, NextColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...predicate.FuelRecord) predicate.FuelRecord {
	return predicate.FuelRecord(func(s *sql.Selector) {
		step := newNextStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() predicate.FuelRecord {
	return predicate.FuelRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...predicate.FuelRecord) predicate.FuelRecord {
	return predicate.FuelRecord(func(s *sql.Selector) {
		step := newPrevStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FuelRecord) predicate.FuelRecord {
	return predicate.FuelRecord(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FuelRecord) predicate.FuelRecord {
	return predicate.FuelRecord(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FuelRecord) predicate.FuelRecord {
	return predicate.FuelRecord(sql.NotPredicates(p))
}
