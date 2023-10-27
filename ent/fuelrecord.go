// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/car"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
)

// FuelRecord is the model entity for the FuelRecord schema.
type FuelRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CurrentFuelLiters holds the value of the "current_fuel_liters" field.
	CurrentFuelLiters float64 `json:"current_fuel_liters,omitempty"`
	// Difference holds the value of the "difference" field.
	Difference float64 `json:"difference,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CarID holds the value of the "car_id" field.
	CarID uuid.UUID `json:"car_id,omitempty"`
	// NextFuelRecordID holds the value of the "next_fuel_record_id" field.
	NextFuelRecordID *uuid.UUID `json:"next_fuel_record_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FuelRecordQuery when eager-loading is set.
	Edges        FuelRecordEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FuelRecordEdges holds the relations/edges for other nodes in the graph.
type FuelRecordEdges struct {
	// Car holds the value of the car edge.
	Car *Car `json:"car,omitempty"`
	// Next holds the value of the next edge.
	Next *FuelRecord `json:"next,omitempty"`
	// Prev holds the value of the prev edge.
	Prev *FuelRecord `json:"prev,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CarOrErr returns the Car value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FuelRecordEdges) CarOrErr() (*Car, error) {
	if e.loadedTypes[0] {
		if e.Car == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: car.Label}
		}
		return e.Car, nil
	}
	return nil, &NotLoadedError{edge: "car"}
}

// NextOrErr returns the Next value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FuelRecordEdges) NextOrErr() (*FuelRecord, error) {
	if e.loadedTypes[1] {
		if e.Next == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: fuelrecord.Label}
		}
		return e.Next, nil
	}
	return nil, &NotLoadedError{edge: "next"}
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FuelRecordEdges) PrevOrErr() (*FuelRecord, error) {
	if e.loadedTypes[2] {
		if e.Prev == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: fuelrecord.Label}
		}
		return e.Prev, nil
	}
	return nil, &NotLoadedError{edge: "prev"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FuelRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fuelrecord.FieldNextFuelRecordID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case fuelrecord.FieldCurrentFuelLiters, fuelrecord.FieldDifference:
			values[i] = new(sql.NullFloat64)
		case fuelrecord.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case fuelrecord.FieldID, fuelrecord.FieldCarID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FuelRecord fields.
func (fr *FuelRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fuelrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fr.ID = *value
			}
		case fuelrecord.FieldCurrentFuelLiters:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field current_fuel_liters", values[i])
			} else if value.Valid {
				fr.CurrentFuelLiters = value.Float64
			}
		case fuelrecord.FieldDifference:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field difference", values[i])
			} else if value.Valid {
				fr.Difference = value.Float64
			}
		case fuelrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fr.CreatedAt = value.Time
			}
		case fuelrecord.FieldCarID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field car_id", values[i])
			} else if value != nil {
				fr.CarID = *value
			}
		case fuelrecord.FieldNextFuelRecordID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field next_fuel_record_id", values[i])
			} else if value.Valid {
				fr.NextFuelRecordID = new(uuid.UUID)
				*fr.NextFuelRecordID = *value.S.(*uuid.UUID)
			}
		default:
			fr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FuelRecord.
// This includes values selected through modifiers, order, etc.
func (fr *FuelRecord) Value(name string) (ent.Value, error) {
	return fr.selectValues.Get(name)
}

// QueryCar queries the "car" edge of the FuelRecord entity.
func (fr *FuelRecord) QueryCar() *CarQuery {
	return NewFuelRecordClient(fr.config).QueryCar(fr)
}

// QueryNext queries the "next" edge of the FuelRecord entity.
func (fr *FuelRecord) QueryNext() *FuelRecordQuery {
	return NewFuelRecordClient(fr.config).QueryNext(fr)
}

// QueryPrev queries the "prev" edge of the FuelRecord entity.
func (fr *FuelRecord) QueryPrev() *FuelRecordQuery {
	return NewFuelRecordClient(fr.config).QueryPrev(fr)
}

// Update returns a builder for updating this FuelRecord.
// Note that you need to call FuelRecord.Unwrap() before calling this method if this FuelRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FuelRecord) Update() *FuelRecordUpdateOne {
	return NewFuelRecordClient(fr.config).UpdateOne(fr)
}

// Unwrap unwraps the FuelRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FuelRecord) Unwrap() *FuelRecord {
	_tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FuelRecord is not a transactional entity")
	}
	fr.config.driver = _tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FuelRecord) String() string {
	var builder strings.Builder
	builder.WriteString("FuelRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fr.ID))
	builder.WriteString("current_fuel_liters=")
	builder.WriteString(fmt.Sprintf("%v", fr.CurrentFuelLiters))
	builder.WriteString(", ")
	builder.WriteString("difference=")
	builder.WriteString(fmt.Sprintf("%v", fr.Difference))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("car_id=")
	builder.WriteString(fmt.Sprintf("%v", fr.CarID))
	builder.WriteString(", ")
	if v := fr.NextFuelRecordID; v != nil {
		builder.WriteString("next_fuel_record_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// FuelRecords is a parsable slice of FuelRecord.
type FuelRecords []*FuelRecord
