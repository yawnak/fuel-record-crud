// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "car_id", Type: field.TypeUUID},
		{Name: "make", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "year", Type: field.TypeInt32},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
	}
	// FuelRecordsColumns holds the columns for the "fuel_records" table.
	FuelRecordsColumns = []*schema.Column{
		{Name: "fuel_record_id", Type: field.TypeUUID},
		{Name: "current_fuel_liters", Type: field.TypeFloat64},
		{Name: "difference", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "car_id", Type: field.TypeUUID},
		{Name: "next_fuel_record_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// FuelRecordsTable holds the schema information for the "fuel_records" table.
	FuelRecordsTable = &schema.Table{
		Name:       "fuel_records",
		Columns:    FuelRecordsColumns,
		PrimaryKey: []*schema.Column{FuelRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "fuel_records_cars_fuel_records",
				Columns:    []*schema.Column{FuelRecordsColumns[4]},
				RefColumns: []*schema.Column{CarsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "fuel_records_fuel_records_prev",
				Columns:    []*schema.Column{FuelRecordsColumns[5]},
				RefColumns: []*schema.Column{FuelRecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OdometerRecordsColumns holds the columns for the "odometer_records" table.
	OdometerRecordsColumns = []*schema.Column{
		{Name: "odometer_record_id", Type: field.TypeUUID},
		{Name: "current_kilometers", Type: field.TypeFloat64},
		{Name: "difference", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "car_id", Type: field.TypeUUID},
		{Name: "next_odometer_record_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// OdometerRecordsTable holds the schema information for the "odometer_records" table.
	OdometerRecordsTable = &schema.Table{
		Name:       "odometer_records",
		Columns:    OdometerRecordsColumns,
		PrimaryKey: []*schema.Column{OdometerRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "odometer_records_cars_odometer_records",
				Columns:    []*schema.Column{OdometerRecordsColumns[4]},
				RefColumns: []*schema.Column{CarsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "odometer_records_odometer_records_prev",
				Columns:    []*schema.Column{OdometerRecordsColumns[5]},
				RefColumns: []*schema.Column{OdometerRecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		FuelRecordsTable,
		OdometerRecordsTable,
	}
)

func init() {
	FuelRecordsTable.ForeignKeys[0].RefTable = CarsTable
	FuelRecordsTable.ForeignKeys[1].RefTable = FuelRecordsTable
	OdometerRecordsTable.ForeignKeys[0].RefTable = CarsTable
	OdometerRecordsTable.ForeignKeys[1].RefTable = OdometerRecordsTable
}
