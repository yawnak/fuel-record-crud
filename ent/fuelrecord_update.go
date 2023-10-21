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
	"github.com/yawnak/fuel-record-crud/ent/predicate"
)

// FuelRecordUpdate is the builder for updating FuelRecord entities.
type FuelRecordUpdate struct {
	config
	hooks    []Hook
	mutation *FuelRecordMutation
}

// Where appends a list predicates to the FuelRecordUpdate builder.
func (fru *FuelRecordUpdate) Where(ps ...predicate.FuelRecord) *FuelRecordUpdate {
	fru.mutation.Where(ps...)
	return fru
}

// SetCurrentFuelLiters sets the "current_fuel_liters" field.
func (fru *FuelRecordUpdate) SetCurrentFuelLiters(f float64) *FuelRecordUpdate {
	fru.mutation.ResetCurrentFuelLiters()
	fru.mutation.SetCurrentFuelLiters(f)
	return fru
}

// AddCurrentFuelLiters adds f to the "current_fuel_liters" field.
func (fru *FuelRecordUpdate) AddCurrentFuelLiters(f float64) *FuelRecordUpdate {
	fru.mutation.AddCurrentFuelLiters(f)
	return fru
}

// SetDifference sets the "difference" field.
func (fru *FuelRecordUpdate) SetDifference(f float64) *FuelRecordUpdate {
	fru.mutation.ResetDifference()
	fru.mutation.SetDifference(f)
	return fru
}

// AddDifference adds f to the "difference" field.
func (fru *FuelRecordUpdate) AddDifference(f float64) *FuelRecordUpdate {
	fru.mutation.AddDifference(f)
	return fru
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (fru *FuelRecordUpdate) SetCarID(id uuid.UUID) *FuelRecordUpdate {
	fru.mutation.SetCarID(id)
	return fru
}

// SetCar sets the "car" edge to the Car entity.
func (fru *FuelRecordUpdate) SetCar(c *Car) *FuelRecordUpdate {
	return fru.SetCarID(c.ID)
}

// SetNextID sets the "next" edge to the FuelRecord entity by ID.
func (fru *FuelRecordUpdate) SetNextID(id uuid.UUID) *FuelRecordUpdate {
	fru.mutation.SetNextID(id)
	return fru
}

// SetNillableNextID sets the "next" edge to the FuelRecord entity by ID if the given value is not nil.
func (fru *FuelRecordUpdate) SetNillableNextID(id *uuid.UUID) *FuelRecordUpdate {
	if id != nil {
		fru = fru.SetNextID(*id)
	}
	return fru
}

// SetNext sets the "next" edge to the FuelRecord entity.
func (fru *FuelRecordUpdate) SetNext(f *FuelRecord) *FuelRecordUpdate {
	return fru.SetNextID(f.ID)
}

// SetPrevID sets the "prev" edge to the FuelRecord entity by ID.
func (fru *FuelRecordUpdate) SetPrevID(id uuid.UUID) *FuelRecordUpdate {
	fru.mutation.SetPrevID(id)
	return fru
}

// SetNillablePrevID sets the "prev" edge to the FuelRecord entity by ID if the given value is not nil.
func (fru *FuelRecordUpdate) SetNillablePrevID(id *uuid.UUID) *FuelRecordUpdate {
	if id != nil {
		fru = fru.SetPrevID(*id)
	}
	return fru
}

// SetPrev sets the "prev" edge to the FuelRecord entity.
func (fru *FuelRecordUpdate) SetPrev(f *FuelRecord) *FuelRecordUpdate {
	return fru.SetPrevID(f.ID)
}

// Mutation returns the FuelRecordMutation object of the builder.
func (fru *FuelRecordUpdate) Mutation() *FuelRecordMutation {
	return fru.mutation
}

// ClearCar clears the "car" edge to the Car entity.
func (fru *FuelRecordUpdate) ClearCar() *FuelRecordUpdate {
	fru.mutation.ClearCar()
	return fru
}

// ClearNext clears the "next" edge to the FuelRecord entity.
func (fru *FuelRecordUpdate) ClearNext() *FuelRecordUpdate {
	fru.mutation.ClearNext()
	return fru
}

// ClearPrev clears the "prev" edge to the FuelRecord entity.
func (fru *FuelRecordUpdate) ClearPrev() *FuelRecordUpdate {
	fru.mutation.ClearPrev()
	return fru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fru *FuelRecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fru.sqlSave, fru.mutation, fru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fru *FuelRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := fru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fru *FuelRecordUpdate) Exec(ctx context.Context) error {
	_, err := fru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fru *FuelRecordUpdate) ExecX(ctx context.Context) {
	if err := fru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fru *FuelRecordUpdate) check() error {
	if v, ok := fru.mutation.CurrentFuelLiters(); ok {
		if err := fuelrecord.CurrentFuelLitersValidator(v); err != nil {
			return &ValidationError{Name: "current_fuel_liters", err: fmt.Errorf(`ent: validator failed for field "FuelRecord.current_fuel_liters": %w`, err)}
		}
	}
	if _, ok := fru.mutation.CarID(); fru.mutation.CarCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FuelRecord.car"`)
	}
	return nil
}

func (fru *FuelRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(fuelrecord.Table, fuelrecord.Columns, sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID))
	if ps := fru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fru.mutation.CurrentFuelLiters(); ok {
		_spec.SetField(fuelrecord.FieldCurrentFuelLiters, field.TypeFloat64, value)
	}
	if value, ok := fru.mutation.AddedCurrentFuelLiters(); ok {
		_spec.AddField(fuelrecord.FieldCurrentFuelLiters, field.TypeFloat64, value)
	}
	if value, ok := fru.mutation.Difference(); ok {
		_spec.SetField(fuelrecord.FieldDifference, field.TypeFloat64, value)
	}
	if value, ok := fru.mutation.AddedDifference(); ok {
		_spec.AddField(fuelrecord.FieldDifference, field.TypeFloat64, value)
	}
	if fru.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fuelrecord.CarTable,
			Columns: []string{fuelrecord.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fuelrecord.CarTable,
			Columns: []string{fuelrecord.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fru.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   fuelrecord.NextTable,
			Columns: []string{fuelrecord.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   fuelrecord.NextTable,
			Columns: []string{fuelrecord.NextColumn},
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
	if fru.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   fuelrecord.PrevTable,
			Columns: []string{fuelrecord.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   fuelrecord.PrevTable,
			Columns: []string{fuelrecord.PrevColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fuelrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fru.mutation.done = true
	return n, nil
}

// FuelRecordUpdateOne is the builder for updating a single FuelRecord entity.
type FuelRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FuelRecordMutation
}

// SetCurrentFuelLiters sets the "current_fuel_liters" field.
func (fruo *FuelRecordUpdateOne) SetCurrentFuelLiters(f float64) *FuelRecordUpdateOne {
	fruo.mutation.ResetCurrentFuelLiters()
	fruo.mutation.SetCurrentFuelLiters(f)
	return fruo
}

// AddCurrentFuelLiters adds f to the "current_fuel_liters" field.
func (fruo *FuelRecordUpdateOne) AddCurrentFuelLiters(f float64) *FuelRecordUpdateOne {
	fruo.mutation.AddCurrentFuelLiters(f)
	return fruo
}

// SetDifference sets the "difference" field.
func (fruo *FuelRecordUpdateOne) SetDifference(f float64) *FuelRecordUpdateOne {
	fruo.mutation.ResetDifference()
	fruo.mutation.SetDifference(f)
	return fruo
}

// AddDifference adds f to the "difference" field.
func (fruo *FuelRecordUpdateOne) AddDifference(f float64) *FuelRecordUpdateOne {
	fruo.mutation.AddDifference(f)
	return fruo
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (fruo *FuelRecordUpdateOne) SetCarID(id uuid.UUID) *FuelRecordUpdateOne {
	fruo.mutation.SetCarID(id)
	return fruo
}

// SetCar sets the "car" edge to the Car entity.
func (fruo *FuelRecordUpdateOne) SetCar(c *Car) *FuelRecordUpdateOne {
	return fruo.SetCarID(c.ID)
}

// SetNextID sets the "next" edge to the FuelRecord entity by ID.
func (fruo *FuelRecordUpdateOne) SetNextID(id uuid.UUID) *FuelRecordUpdateOne {
	fruo.mutation.SetNextID(id)
	return fruo
}

// SetNillableNextID sets the "next" edge to the FuelRecord entity by ID if the given value is not nil.
func (fruo *FuelRecordUpdateOne) SetNillableNextID(id *uuid.UUID) *FuelRecordUpdateOne {
	if id != nil {
		fruo = fruo.SetNextID(*id)
	}
	return fruo
}

// SetNext sets the "next" edge to the FuelRecord entity.
func (fruo *FuelRecordUpdateOne) SetNext(f *FuelRecord) *FuelRecordUpdateOne {
	return fruo.SetNextID(f.ID)
}

// SetPrevID sets the "prev" edge to the FuelRecord entity by ID.
func (fruo *FuelRecordUpdateOne) SetPrevID(id uuid.UUID) *FuelRecordUpdateOne {
	fruo.mutation.SetPrevID(id)
	return fruo
}

// SetNillablePrevID sets the "prev" edge to the FuelRecord entity by ID if the given value is not nil.
func (fruo *FuelRecordUpdateOne) SetNillablePrevID(id *uuid.UUID) *FuelRecordUpdateOne {
	if id != nil {
		fruo = fruo.SetPrevID(*id)
	}
	return fruo
}

// SetPrev sets the "prev" edge to the FuelRecord entity.
func (fruo *FuelRecordUpdateOne) SetPrev(f *FuelRecord) *FuelRecordUpdateOne {
	return fruo.SetPrevID(f.ID)
}

// Mutation returns the FuelRecordMutation object of the builder.
func (fruo *FuelRecordUpdateOne) Mutation() *FuelRecordMutation {
	return fruo.mutation
}

// ClearCar clears the "car" edge to the Car entity.
func (fruo *FuelRecordUpdateOne) ClearCar() *FuelRecordUpdateOne {
	fruo.mutation.ClearCar()
	return fruo
}

// ClearNext clears the "next" edge to the FuelRecord entity.
func (fruo *FuelRecordUpdateOne) ClearNext() *FuelRecordUpdateOne {
	fruo.mutation.ClearNext()
	return fruo
}

// ClearPrev clears the "prev" edge to the FuelRecord entity.
func (fruo *FuelRecordUpdateOne) ClearPrev() *FuelRecordUpdateOne {
	fruo.mutation.ClearPrev()
	return fruo
}

// Where appends a list predicates to the FuelRecordUpdate builder.
func (fruo *FuelRecordUpdateOne) Where(ps ...predicate.FuelRecord) *FuelRecordUpdateOne {
	fruo.mutation.Where(ps...)
	return fruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fruo *FuelRecordUpdateOne) Select(field string, fields ...string) *FuelRecordUpdateOne {
	fruo.fields = append([]string{field}, fields...)
	return fruo
}

// Save executes the query and returns the updated FuelRecord entity.
func (fruo *FuelRecordUpdateOne) Save(ctx context.Context) (*FuelRecord, error) {
	return withHooks(ctx, fruo.sqlSave, fruo.mutation, fruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fruo *FuelRecordUpdateOne) SaveX(ctx context.Context) *FuelRecord {
	node, err := fruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fruo *FuelRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := fruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fruo *FuelRecordUpdateOne) ExecX(ctx context.Context) {
	if err := fruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fruo *FuelRecordUpdateOne) check() error {
	if v, ok := fruo.mutation.CurrentFuelLiters(); ok {
		if err := fuelrecord.CurrentFuelLitersValidator(v); err != nil {
			return &ValidationError{Name: "current_fuel_liters", err: fmt.Errorf(`ent: validator failed for field "FuelRecord.current_fuel_liters": %w`, err)}
		}
	}
	if _, ok := fruo.mutation.CarID(); fruo.mutation.CarCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FuelRecord.car"`)
	}
	return nil
}

func (fruo *FuelRecordUpdateOne) sqlSave(ctx context.Context) (_node *FuelRecord, err error) {
	if err := fruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(fuelrecord.Table, fuelrecord.Columns, sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID))
	id, ok := fruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FuelRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fuelrecord.FieldID)
		for _, f := range fields {
			if !fuelrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fuelrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fruo.mutation.CurrentFuelLiters(); ok {
		_spec.SetField(fuelrecord.FieldCurrentFuelLiters, field.TypeFloat64, value)
	}
	if value, ok := fruo.mutation.AddedCurrentFuelLiters(); ok {
		_spec.AddField(fuelrecord.FieldCurrentFuelLiters, field.TypeFloat64, value)
	}
	if value, ok := fruo.mutation.Difference(); ok {
		_spec.SetField(fuelrecord.FieldDifference, field.TypeFloat64, value)
	}
	if value, ok := fruo.mutation.AddedDifference(); ok {
		_spec.AddField(fuelrecord.FieldDifference, field.TypeFloat64, value)
	}
	if fruo.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fuelrecord.CarTable,
			Columns: []string{fuelrecord.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fuelrecord.CarTable,
			Columns: []string{fuelrecord.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fruo.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   fuelrecord.NextTable,
			Columns: []string{fuelrecord.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   fuelrecord.NextTable,
			Columns: []string{fuelrecord.NextColumn},
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
	if fruo.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   fuelrecord.PrevTable,
			Columns: []string{fuelrecord.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   fuelrecord.PrevTable,
			Columns: []string{fuelrecord.PrevColumn},
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
	_node = &FuelRecord{config: fruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fuelrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fruo.mutation.done = true
	return _node, nil
}
