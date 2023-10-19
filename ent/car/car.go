// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "car_id"
	// FieldMake holds the string denoting the make field in the database.
	FieldMake = "make"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// EdgeFuelRecords holds the string denoting the fuel_records edge name in mutations.
	EdgeFuelRecords = "fuel_records"
	// EdgeOdometerRecords holds the string denoting the odometer_records edge name in mutations.
	EdgeOdometerRecords = "odometer_records"
	// FuelRecordFieldID holds the string denoting the ID field of the FuelRecord.
	FuelRecordFieldID = "fuel_record_id"
	// OdometerRecordFieldID holds the string denoting the ID field of the OdometerRecord.
	OdometerRecordFieldID = "odometer_record_id"
	// Table holds the table name of the car in the database.
	Table = "cars"
	// FuelRecordsTable is the table that holds the fuel_records relation/edge.
	FuelRecordsTable = "fuel_records"
	// FuelRecordsInverseTable is the table name for the FuelRecord entity.
	// It exists in this package in order to avoid circular dependency with the "fuelrecord" package.
	FuelRecordsInverseTable = "fuel_records"
	// FuelRecordsColumn is the table column denoting the fuel_records relation/edge.
	FuelRecordsColumn = "car_fuel_records"
	// OdometerRecordsTable is the table that holds the odometer_records relation/edge.
	OdometerRecordsTable = "odometer_records"
	// OdometerRecordsInverseTable is the table name for the OdometerRecord entity.
	// It exists in this package in order to avoid circular dependency with the "odometerrecord" package.
	OdometerRecordsInverseTable = "odometer_records"
	// OdometerRecordsColumn is the table column denoting the odometer_records relation/edge.
	OdometerRecordsColumn = "car_odometer_records"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldMake,
	FieldModel,
	FieldYear,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// MakeValidator is a validator for the "make" field. It is called by the builders before save.
	MakeValidator func(string) error
	// ModelValidator is a validator for the "model" field. It is called by the builders before save.
	ModelValidator func(string) error
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int32) error
)

// OrderOption defines the ordering options for the Car queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMake orders the results by the make field.
func ByMake(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMake, opts...).ToFunc()
}

// ByModel orders the results by the model field.
func ByModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModel, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByFuelRecordsCount orders the results by fuel_records count.
func ByFuelRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFuelRecordsStep(), opts...)
	}
}

// ByFuelRecords orders the results by fuel_records terms.
func ByFuelRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFuelRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOdometerRecordsCount orders the results by odometer_records count.
func ByOdometerRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOdometerRecordsStep(), opts...)
	}
}

// ByOdometerRecords orders the results by odometer_records terms.
func ByOdometerRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOdometerRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFuelRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FuelRecordsInverseTable, FuelRecordFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FuelRecordsTable, FuelRecordsColumn),
	)
}
func newOdometerRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OdometerRecordsInverseTable, OdometerRecordFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OdometerRecordsTable, OdometerRecordsColumn),
	)
}
