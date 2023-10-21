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
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
)

// OdometerRecord is the model entity for the OdometerRecord schema.
type OdometerRecord struct {
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
	// NextOdometerRecordID holds the value of the "next_odometer_record_id" field.
	NextOdometerRecordID uuid.UUID `json:"next_odometer_record_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OdometerRecordQuery when eager-loading is set.
	Edges        OdometerRecordEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OdometerRecordEdges holds the relations/edges for other nodes in the graph.
type OdometerRecordEdges struct {
	// Car holds the value of the car edge.
	Car *Car `json:"car,omitempty"`
	// Prev holds the value of the prev edge.
	Prev *OdometerRecord `json:"prev,omitempty"`
	// Next holds the value of the next edge.
	Next *OdometerRecord `json:"next,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CarOrErr returns the Car value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OdometerRecordEdges) CarOrErr() (*Car, error) {
	if e.loadedTypes[0] {
		if e.Car == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: car.Label}
		}
		return e.Car, nil
	}
	return nil, &NotLoadedError{edge: "car"}
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OdometerRecordEdges) PrevOrErr() (*OdometerRecord, error) {
	if e.loadedTypes[1] {
		if e.Prev == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: odometerrecord.Label}
		}
		return e.Prev, nil
	}
	return nil, &NotLoadedError{edge: "prev"}
}

// NextOrErr returns the Next value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OdometerRecordEdges) NextOrErr() (*OdometerRecord, error) {
	if e.loadedTypes[2] {
		if e.Next == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: odometerrecord.Label}
		}
		return e.Next, nil
	}
	return nil, &NotLoadedError{edge: "next"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OdometerRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case odometerrecord.FieldCurrentFuelLiters, odometerrecord.FieldDifference:
			values[i] = new(sql.NullFloat64)
		case odometerrecord.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case odometerrecord.FieldID, odometerrecord.FieldCarID, odometerrecord.FieldNextOdometerRecordID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OdometerRecord fields.
func (or *OdometerRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case odometerrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				or.ID = *value
			}
		case odometerrecord.FieldCurrentFuelLiters:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field current_fuel_liters", values[i])
			} else if value.Valid {
				or.CurrentFuelLiters = value.Float64
			}
		case odometerrecord.FieldDifference:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field difference", values[i])
			} else if value.Valid {
				or.Difference = value.Float64
			}
		case odometerrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				or.CreatedAt = value.Time
			}
		case odometerrecord.FieldCarID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field car_id", values[i])
			} else if value != nil {
				or.CarID = *value
			}
		case odometerrecord.FieldNextOdometerRecordID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field next_odometer_record_id", values[i])
			} else if value != nil {
				or.NextOdometerRecordID = *value
			}
		default:
			or.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OdometerRecord.
// This includes values selected through modifiers, order, etc.
func (or *OdometerRecord) Value(name string) (ent.Value, error) {
	return or.selectValues.Get(name)
}

// QueryCar queries the "car" edge of the OdometerRecord entity.
func (or *OdometerRecord) QueryCar() *CarQuery {
	return NewOdometerRecordClient(or.config).QueryCar(or)
}

// QueryPrev queries the "prev" edge of the OdometerRecord entity.
func (or *OdometerRecord) QueryPrev() *OdometerRecordQuery {
	return NewOdometerRecordClient(or.config).QueryPrev(or)
}

// QueryNext queries the "next" edge of the OdometerRecord entity.
func (or *OdometerRecord) QueryNext() *OdometerRecordQuery {
	return NewOdometerRecordClient(or.config).QueryNext(or)
}

// Update returns a builder for updating this OdometerRecord.
// Note that you need to call OdometerRecord.Unwrap() before calling this method if this OdometerRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OdometerRecord) Update() *OdometerRecordUpdateOne {
	return NewOdometerRecordClient(or.config).UpdateOne(or)
}

// Unwrap unwraps the OdometerRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OdometerRecord) Unwrap() *OdometerRecord {
	_tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OdometerRecord is not a transactional entity")
	}
	or.config.driver = _tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OdometerRecord) String() string {
	var builder strings.Builder
	builder.WriteString("OdometerRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", or.ID))
	builder.WriteString("current_fuel_liters=")
	builder.WriteString(fmt.Sprintf("%v", or.CurrentFuelLiters))
	builder.WriteString(", ")
	builder.WriteString("difference=")
	builder.WriteString(fmt.Sprintf("%v", or.Difference))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(or.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("car_id=")
	builder.WriteString(fmt.Sprintf("%v", or.CarID))
	builder.WriteString(", ")
	builder.WriteString("next_odometer_record_id=")
	builder.WriteString(fmt.Sprintf("%v", or.NextOdometerRecordID))
	builder.WriteByte(')')
	return builder.String()
}

// OdometerRecords is a parsable slice of OdometerRecord.
type OdometerRecords []*OdometerRecord
