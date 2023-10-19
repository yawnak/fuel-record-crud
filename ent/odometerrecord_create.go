// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/yawnak/fuel-record-crud/ent/car"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
)

// OdometerRecordCreate is the builder for creating a OdometerRecord entity.
type OdometerRecordCreate struct {
	config
	mutation *OdometerRecordMutation
	hooks    []Hook
}

// SetCurrentFuelLiters sets the "current_fuel_liters" field.
func (orc *OdometerRecordCreate) SetCurrentFuelLiters(f float64) *OdometerRecordCreate {
	orc.mutation.SetCurrentFuelLiters(f)
	return orc
}

// SetDifference sets the "difference" field.
func (orc *OdometerRecordCreate) SetDifference(f float64) *OdometerRecordCreate {
	orc.mutation.SetDifference(f)
	return orc
}

// SetCreatedAt sets the "created_at" field.
func (orc *OdometerRecordCreate) SetCreatedAt(t time.Time) *OdometerRecordCreate {
	orc.mutation.SetCreatedAt(t)
	return orc
}

// SetID sets the "id" field.
func (orc *OdometerRecordCreate) SetID(u uuid.UUID) *OdometerRecordCreate {
	orc.mutation.SetID(u)
	return orc
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (orc *OdometerRecordCreate) SetCarID(id uuid.UUID) *OdometerRecordCreate {
	orc.mutation.SetCarID(id)
	return orc
}

// SetCar sets the "car" edge to the Car entity.
func (orc *OdometerRecordCreate) SetCar(c *Car) *OdometerRecordCreate {
	return orc.SetCarID(c.ID)
}

// SetPrevID sets the "prev" edge to the OdometerRecord entity by ID.
func (orc *OdometerRecordCreate) SetPrevID(id uuid.UUID) *OdometerRecordCreate {
	orc.mutation.SetPrevID(id)
	return orc
}

// SetNillablePrevID sets the "prev" edge to the OdometerRecord entity by ID if the given value is not nil.
func (orc *OdometerRecordCreate) SetNillablePrevID(id *uuid.UUID) *OdometerRecordCreate {
	if id != nil {
		orc = orc.SetPrevID(*id)
	}
	return orc
}

// SetPrev sets the "prev" edge to the OdometerRecord entity.
func (orc *OdometerRecordCreate) SetPrev(o *OdometerRecord) *OdometerRecordCreate {
	return orc.SetPrevID(o.ID)
}

// SetNextID sets the "next" edge to the OdometerRecord entity by ID.
func (orc *OdometerRecordCreate) SetNextID(id uuid.UUID) *OdometerRecordCreate {
	orc.mutation.SetNextID(id)
	return orc
}

// SetNillableNextID sets the "next" edge to the OdometerRecord entity by ID if the given value is not nil.
func (orc *OdometerRecordCreate) SetNillableNextID(id *uuid.UUID) *OdometerRecordCreate {
	if id != nil {
		orc = orc.SetNextID(*id)
	}
	return orc
}

// SetNext sets the "next" edge to the OdometerRecord entity.
func (orc *OdometerRecordCreate) SetNext(o *OdometerRecord) *OdometerRecordCreate {
	return orc.SetNextID(o.ID)
}

// Mutation returns the OdometerRecordMutation object of the builder.
func (orc *OdometerRecordCreate) Mutation() *OdometerRecordMutation {
	return orc.mutation
}

// Save creates the OdometerRecord in the database.
func (orc *OdometerRecordCreate) Save(ctx context.Context) (*OdometerRecord, error) {
	return withHooks(ctx, orc.sqlSave, orc.mutation, orc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (orc *OdometerRecordCreate) SaveX(ctx context.Context) *OdometerRecord {
	v, err := orc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orc *OdometerRecordCreate) Exec(ctx context.Context) error {
	_, err := orc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orc *OdometerRecordCreate) ExecX(ctx context.Context) {
	if err := orc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (orc *OdometerRecordCreate) check() error {
	if _, ok := orc.mutation.CurrentFuelLiters(); !ok {
		return &ValidationError{Name: "current_fuel_liters", err: errors.New(`ent: missing required field "OdometerRecord.current_fuel_liters"`)}
	}
	if v, ok := orc.mutation.CurrentFuelLiters(); ok {
		if err := odometerrecord.CurrentFuelLitersValidator(v); err != nil {
			return &ValidationError{Name: "current_fuel_liters", err: fmt.Errorf(`ent: validator failed for field "OdometerRecord.current_fuel_liters": %w`, err)}
		}
	}
	if _, ok := orc.mutation.Difference(); !ok {
		return &ValidationError{Name: "difference", err: errors.New(`ent: missing required field "OdometerRecord.difference"`)}
	}
	if v, ok := orc.mutation.Difference(); ok {
		if err := odometerrecord.DifferenceValidator(v); err != nil {
			return &ValidationError{Name: "difference", err: fmt.Errorf(`ent: validator failed for field "OdometerRecord.difference": %w`, err)}
		}
	}
	if _, ok := orc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OdometerRecord.created_at"`)}
	}
	if _, ok := orc.mutation.CarID(); !ok {
		return &ValidationError{Name: "car", err: errors.New(`ent: missing required edge "OdometerRecord.car"`)}
	}
	return nil
}

func (orc *OdometerRecordCreate) sqlSave(ctx context.Context) (*OdometerRecord, error) {
	if err := orc.check(); err != nil {
		return nil, err
	}
	_node, _spec := orc.createSpec()
	if err := sqlgraph.CreateNode(ctx, orc.driver, _spec); err != nil {
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
	orc.mutation.id = &_node.ID
	orc.mutation.done = true
	return _node, nil
}

func (orc *OdometerRecordCreate) createSpec() (*OdometerRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &OdometerRecord{config: orc.config}
		_spec = sqlgraph.NewCreateSpec(odometerrecord.Table, sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID))
	)
	if id, ok := orc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := orc.mutation.CurrentFuelLiters(); ok {
		_spec.SetField(odometerrecord.FieldCurrentFuelLiters, field.TypeFloat64, value)
		_node.CurrentFuelLiters = value
	}
	if value, ok := orc.mutation.Difference(); ok {
		_spec.SetField(odometerrecord.FieldDifference, field.TypeFloat64, value)
		_node.Difference = value
	}
	if value, ok := orc.mutation.CreatedAt(); ok {
		_spec.SetField(odometerrecord.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := orc.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   odometerrecord.CarTable,
			Columns: []string{odometerrecord.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.car_odometer_records = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := orc.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   odometerrecord.PrevTable,
			Columns: []string{odometerrecord.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.odometer_record_next = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := orc.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   odometerrecord.NextTable,
			Columns: []string{odometerrecord.NextColumn},
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

// OdometerRecordCreateBulk is the builder for creating many OdometerRecord entities in bulk.
type OdometerRecordCreateBulk struct {
	config
	err      error
	builders []*OdometerRecordCreate
}

// Save creates the OdometerRecord entities in the database.
func (orcb *OdometerRecordCreateBulk) Save(ctx context.Context) ([]*OdometerRecord, error) {
	if orcb.err != nil {
		return nil, orcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(orcb.builders))
	nodes := make([]*OdometerRecord, len(orcb.builders))
	mutators := make([]Mutator, len(orcb.builders))
	for i := range orcb.builders {
		func(i int, root context.Context) {
			builder := orcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OdometerRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, orcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, orcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, orcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (orcb *OdometerRecordCreateBulk) SaveX(ctx context.Context) []*OdometerRecord {
	v, err := orcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orcb *OdometerRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := orcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orcb *OdometerRecordCreateBulk) ExecX(ctx context.Context) {
	if err := orcb.Exec(ctx); err != nil {
		panic(err)
	}
}