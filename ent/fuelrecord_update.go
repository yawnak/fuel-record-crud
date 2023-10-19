// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
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

// Mutation returns the FuelRecordMutation object of the builder.
func (fru *FuelRecordUpdate) Mutation() *FuelRecordMutation {
	return fru.mutation
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

func (fru *FuelRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(fuelrecord.Table, fuelrecord.Columns, sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeInt))
	if ps := fru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
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

// Mutation returns the FuelRecordMutation object of the builder.
func (fruo *FuelRecordUpdateOne) Mutation() *FuelRecordMutation {
	return fruo.mutation
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

func (fruo *FuelRecordUpdateOne) sqlSave(ctx context.Context) (_node *FuelRecord, err error) {
	_spec := sqlgraph.NewUpdateSpec(fuelrecord.Table, fuelrecord.Columns, sqlgraph.NewFieldSpec(fuelrecord.FieldID, field.TypeInt))
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
