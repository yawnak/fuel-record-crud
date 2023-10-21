// Code generated by ent, DO NOT EDIT.

package odometerrecord

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the odometerrecord type in the database.
	Label = "odometer_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "odometer_record_id"
	// FieldCurrentFuelLiters holds the string denoting the current_fuel_liters field in the database.
	FieldCurrentFuelLiters = "current_fuel_liters"
	// FieldDifference holds the string denoting the difference field in the database.
	FieldDifference = "difference"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"
	// EdgeNext holds the string denoting the next edge name in mutations.
	EdgeNext = "next"
	// EdgePrev holds the string denoting the prev edge name in mutations.
	EdgePrev = "prev"
	// CarFieldID holds the string denoting the ID field of the Car.
	CarFieldID = "car_id"
	// Table holds the table name of the odometerrecord in the database.
	Table = "odometer_records"
	// CarTable is the table that holds the car relation/edge.
	CarTable = "odometer_records"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "car_odometer_records"
	// NextTable is the table that holds the next relation/edge.
	NextTable = "odometer_records"
	// NextColumn is the table column denoting the next relation/edge.
	NextColumn = "odometer_record_prev"
	// PrevTable is the table that holds the prev relation/edge.
	PrevTable = "odometer_records"
	// PrevColumn is the table column denoting the prev relation/edge.
	PrevColumn = "odometer_record_prev"
)

// Columns holds all SQL columns for odometerrecord fields.
var Columns = []string{
	FieldID,
	FieldCurrentFuelLiters,
	FieldDifference,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "odometer_records"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"car_odometer_records",
	"odometer_record_prev",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// CurrentFuelLitersValidator is a validator for the "current_fuel_liters" field. It is called by the builders before save.
	CurrentFuelLitersValidator func(float64) error
	// DifferenceValidator is a validator for the "difference" field. It is called by the builders before save.
	DifferenceValidator func(float64) error
)

// OrderOption defines the ordering options for the OdometerRecord queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCurrentFuelLiters orders the results by the current_fuel_liters field.
func ByCurrentFuelLiters(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentFuelLiters, opts...).ToFunc()
}

// ByDifference orders the results by the difference field.
func ByDifference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDifference, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByCarField orders the results by car field.
func ByCarField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCarStep(), sql.OrderByField(field, opts...))
	}
}

// ByNextField orders the results by next field.
func ByNextField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNextStep(), sql.OrderByField(field, opts...))
	}
}

// ByPrevField orders the results by prev field.
func ByPrevField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrevStep(), sql.OrderByField(field, opts...))
	}
}
func newCarStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CarInverseTable, CarFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CarTable, CarColumn),
	)
}
func newNextStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, NextTable, NextColumn),
	)
}
func newPrevStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, PrevTable, PrevColumn),
	)
}
