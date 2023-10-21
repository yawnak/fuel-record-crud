// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/car"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
}

// SetMake sets the "make" field.
func (cc *CarCreate) SetMake(s string) *CarCreate {
	cc.mutation.SetMake(s)
	return cc
}

// SetModel sets the "model" field.
func (cc *CarCreate) SetModel(s string) *CarCreate {
	cc.mutation.SetModel(s)
	return cc
}

// SetYear sets the "year" field.
func (cc *CarCreate) SetYear(i int32) *CarCreate {
	cc.mutation.SetYear(i)
	return cc
}

// SetID sets the "id" field.
func (cc *CarCreate) SetID(u uuid.UUID) *CarCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddFuelRecordIDs adds the "fuel_records" edge to the FuelRecord entity by IDs.
func (cc *CarCreate) AddFuelRecordIDs(ids ...uuid.UUID) *CarCreate {
	cc.mutation.AddFuelRecordIDs(ids...)
	return cc
}

// AddFuelRecords adds the "fuel_records" edges to the FuelRecord entity.
func (cc *CarCreate) AddFuelRecords(f ...*FuelRecord) *CarCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cc.AddFuelRecordIDs(ids...)
}

// AddOdometerRecordIDs adds the "odometer_records" edge to the OdometerRecord entity by IDs.
func (cc *CarCreate) AddOdometerRecordIDs(ids ...uuid.UUID) *CarCreate {
	cc.mutation.AddOdometerRecordIDs(ids...)
	return cc
}

// AddOdometerRecords adds the "odometer_records" edges to the OdometerRecord entity.
func (cc *CarCreate) AddOdometerRecords(o ...*OdometerRecord) *CarCreate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cc.AddOdometerRecordIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (cc *CarCreate) Mutation() *CarMutation {
	return cc.mutation
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CarCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CarCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CarCreate) check() error {
	if _, ok := cc.mutation.Make(); !ok {
		return &ValidationError{Name: "make", err: errors.New(`ent: missing required field "Car.make"`)}
	}
	if v, ok := cc.mutation.Make(); ok {
		if err := car.MakeValidator(v); err != nil {
			return &ValidationError{Name: "make", err: fmt.Errorf(`ent: validator failed for field "Car.make": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Car.model"`)}
	}
	if v, ok := cc.mutation.Model(); ok {
		if err := car.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "Car.model": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Car.year"`)}
	}
	return nil
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CarCreate) createSpec() (*Car, *sqlgraph.CreateSpec) {
	var (
		_node = &Car{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(car.Table, sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Make(); ok {
		_spec.SetField(car.FieldMake, field.TypeString, value)
		_node.Make = value
	}
	if value, ok := cc.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := cc.mutation.Year(); ok {
		_spec.SetField(car.FieldYear, field.TypeInt32, value)
		_node.Year = value
	}
	if nodes := cc.mutation.FuelRecordsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.OdometerRecordsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CarCreateBulk is the builder for creating many Car entities in bulk.
type CarCreateBulk struct {
	config
	err      error
	builders []*CarCreate
}

// Save creates the Car entities in the database.
func (ccb *CarCreateBulk) Save(ctx context.Context) ([]*Car, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Car, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CarCreateBulk) SaveX(ctx context.Context) []*Car {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CarCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CarCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
