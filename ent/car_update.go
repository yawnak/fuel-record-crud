// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/car"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
	"github.com/yawnak/fuel-record-crud/ent/predicate"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetMake sets the "make" field.
func (cu *CarUpdate) SetMake(s string) *CarUpdate {
	cu.mutation.SetMake(s)
	return cu
}

// SetModel sets the "model" field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.mutation.SetModel(s)
	return cu
}

// SetYear sets the "year" field.
func (cu *CarUpdate) SetYear(i int32) *CarUpdate {
	cu.mutation.ResetYear()
	cu.mutation.SetYear(i)
	return cu
}

// AddYear adds i to the "year" field.
func (cu *CarUpdate) AddYear(i int32) *CarUpdate {
	cu.mutation.AddYear(i)
	return cu
}

// AddFuelRecordIDs adds the "fuel_records" edge to the FuelRecord entity by IDs.
func (cu *CarUpdate) AddFuelRecordIDs(ids ...uuid.UUID) *CarUpdate {
	cu.mutation.AddFuelRecordIDs(ids...)
	return cu
}

// AddFuelRecords adds the "fuel_records" edges to the FuelRecord entity.
func (cu *CarUpdate) AddFuelRecords(f ...*FuelRecord) *CarUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cu.AddFuelRecordIDs(ids...)
}

// AddOdometerRecordIDs adds the "odometer_records" edge to the OdometerRecord entity by IDs.
func (cu *CarUpdate) AddOdometerRecordIDs(ids ...uuid.UUID) *CarUpdate {
	cu.mutation.AddOdometerRecordIDs(ids...)
	return cu
}

// AddOdometerRecords adds the "odometer_records" edges to the OdometerRecord entity.
func (cu *CarUpdate) AddOdometerRecords(o ...*OdometerRecord) *CarUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.AddOdometerRecordIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// ClearFuelRecords clears all "fuel_records" edges to the FuelRecord entity.
func (cu *CarUpdate) ClearFuelRecords() *CarUpdate {
	cu.mutation.ClearFuelRecords()
	return cu
}

// RemoveFuelRecordIDs removes the "fuel_records" edge to FuelRecord entities by IDs.
func (cu *CarUpdate) RemoveFuelRecordIDs(ids ...uuid.UUID) *CarUpdate {
	cu.mutation.RemoveFuelRecordIDs(ids...)
	return cu
}

// RemoveFuelRecords removes "fuel_records" edges to FuelRecord entities.
func (cu *CarUpdate) RemoveFuelRecords(f ...*FuelRecord) *CarUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cu.RemoveFuelRecordIDs(ids...)
}

// ClearOdometerRecords clears all "odometer_records" edges to the OdometerRecord entity.
func (cu *CarUpdate) ClearOdometerRecords() *CarUpdate {
	cu.mutation.ClearOdometerRecords()
	return cu
}

// RemoveOdometerRecordIDs removes the "odometer_records" edge to OdometerRecord entities by IDs.
func (cu *CarUpdate) RemoveOdometerRecordIDs(ids ...uuid.UUID) *CarUpdate {
	cu.mutation.RemoveOdometerRecordIDs(ids...)
	return cu
}

// RemoveOdometerRecords removes "odometer_records" edges to OdometerRecord entities.
func (cu *CarUpdate) RemoveOdometerRecords(o ...*OdometerRecord) *CarUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.RemoveOdometerRecordIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CarUpdate) check() error {
	if v, ok := cu.mutation.Make(); ok {
		if err := car.MakeValidator(v); err != nil {
			return &ValidationError{Name: "make", err: fmt.Errorf(`ent: validator failed for field "Car.make": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Model(); ok {
		if err := car.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "Car.model": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Year(); ok {
		if err := car.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Car.year": %w`, err)}
		}
	}
	return nil
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Make(); ok {
		_spec.SetField(car.FieldMake, field.TypeString, value)
	}
	if value, ok := cu.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if value, ok := cu.mutation.Year(); ok {
		_spec.SetField(car.FieldYear, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedYear(); ok {
		_spec.AddField(car.FieldYear, field.TypeInt32, value)
	}
	if cu.mutation.FuelRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.FuelRecordsTable,
			Columns: []string{car.FuelRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedFuelRecordsIDs(); len(nodes) > 0 && !cu.mutation.FuelRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.FuelRecordsTable,
			Columns: []string{car.FuelRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FuelRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.FuelRecordsTable,
			Columns: []string{car.FuelRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.OdometerRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.OdometerRecordsTable,
			Columns: []string{car.OdometerRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedOdometerRecordsIDs(); len(nodes) > 0 && !cu.mutation.OdometerRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.OdometerRecordsTable,
			Columns: []string{car.OdometerRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OdometerRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.OdometerRecordsTable,
			Columns: []string{car.OdometerRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetMake sets the "make" field.
func (cuo *CarUpdateOne) SetMake(s string) *CarUpdateOne {
	cuo.mutation.SetMake(s)
	return cuo
}

// SetModel sets the "model" field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.mutation.SetModel(s)
	return cuo
}

// SetYear sets the "year" field.
func (cuo *CarUpdateOne) SetYear(i int32) *CarUpdateOne {
	cuo.mutation.ResetYear()
	cuo.mutation.SetYear(i)
	return cuo
}

// AddYear adds i to the "year" field.
func (cuo *CarUpdateOne) AddYear(i int32) *CarUpdateOne {
	cuo.mutation.AddYear(i)
	return cuo
}

// AddFuelRecordIDs adds the "fuel_records" edge to the FuelRecord entity by IDs.
func (cuo *CarUpdateOne) AddFuelRecordIDs(ids ...uuid.UUID) *CarUpdateOne {
	cuo.mutation.AddFuelRecordIDs(ids...)
	return cuo
}

// AddFuelRecords adds the "fuel_records" edges to the FuelRecord entity.
func (cuo *CarUpdateOne) AddFuelRecords(f ...*FuelRecord) *CarUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cuo.AddFuelRecordIDs(ids...)
}

// AddOdometerRecordIDs adds the "odometer_records" edge to the OdometerRecord entity by IDs.
func (cuo *CarUpdateOne) AddOdometerRecordIDs(ids ...uuid.UUID) *CarUpdateOne {
	cuo.mutation.AddOdometerRecordIDs(ids...)
	return cuo
}

// AddOdometerRecords adds the "odometer_records" edges to the OdometerRecord entity.
func (cuo *CarUpdateOne) AddOdometerRecords(o ...*OdometerRecord) *CarUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.AddOdometerRecordIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// ClearFuelRecords clears all "fuel_records" edges to the FuelRecord entity.
func (cuo *CarUpdateOne) ClearFuelRecords() *CarUpdateOne {
	cuo.mutation.ClearFuelRecords()
	return cuo
}

// RemoveFuelRecordIDs removes the "fuel_records" edge to FuelRecord entities by IDs.
func (cuo *CarUpdateOne) RemoveFuelRecordIDs(ids ...uuid.UUID) *CarUpdateOne {
	cuo.mutation.RemoveFuelRecordIDs(ids...)
	return cuo
}

// RemoveFuelRecords removes "fuel_records" edges to FuelRecord entities.
func (cuo *CarUpdateOne) RemoveFuelRecords(f ...*FuelRecord) *CarUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cuo.RemoveFuelRecordIDs(ids...)
}

// ClearOdometerRecords clears all "odometer_records" edges to the OdometerRecord entity.
func (cuo *CarUpdateOne) ClearOdometerRecords() *CarUpdateOne {
	cuo.mutation.ClearOdometerRecords()
	return cuo
}

// RemoveOdometerRecordIDs removes the "odometer_records" edge to OdometerRecord entities by IDs.
func (cuo *CarUpdateOne) RemoveOdometerRecordIDs(ids ...uuid.UUID) *CarUpdateOne {
	cuo.mutation.RemoveOdometerRecordIDs(ids...)
	return cuo
}

// RemoveOdometerRecords removes "odometer_records" edges to OdometerRecord entities.
func (cuo *CarUpdateOne) RemoveOdometerRecords(o ...*OdometerRecord) *CarUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.RemoveOdometerRecordIDs(ids...)
}

// Where appends a list predicates to the CarUpdate builder.
func (cuo *CarUpdateOne) Where(ps ...predicate.Car) *CarUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CarUpdateOne) check() error {
	if v, ok := cuo.mutation.Make(); ok {
		if err := car.MakeValidator(v); err != nil {
			return &ValidationError{Name: "make", err: fmt.Errorf(`ent: validator failed for field "Car.make": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Model(); ok {
		if err := car.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "Car.model": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Year(); ok {
		if err := car.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Car.year": %w`, err)}
		}
	}
	return nil
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Make(); ok {
		_spec.SetField(car.FieldMake, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Year(); ok {
		_spec.SetField(car.FieldYear, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedYear(); ok {
		_spec.AddField(car.FieldYear, field.TypeInt32, value)
	}
	if cuo.mutation.FuelRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.FuelRecordsTable,
			Columns: []string{car.FuelRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedFuelRecordsIDs(); len(nodes) > 0 && !cuo.mutation.FuelRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.FuelRecordsTable,
			Columns: []string{car.FuelRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FuelRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.FuelRecordsTable,
			Columns: []string{car.FuelRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.OdometerRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.OdometerRecordsTable,
			Columns: []string{car.OdometerRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedOdometerRecordsIDs(); len(nodes) > 0 && !cuo.mutation.OdometerRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.OdometerRecordsTable,
			Columns: []string{car.OdometerRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OdometerRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.OdometerRecordsTable,
			Columns: []string{car.OdometerRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Car{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
