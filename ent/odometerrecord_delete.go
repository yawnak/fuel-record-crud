// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
	"github.com/yawnak/fuel-record-crud/ent/predicate"
)

// OdometerRecordDelete is the builder for deleting a OdometerRecord entity.
type OdometerRecordDelete struct {
	config
	hooks    []Hook
	mutation *OdometerRecordMutation
}

// Where appends a list predicates to the OdometerRecordDelete builder.
func (ord *OdometerRecordDelete) Where(ps ...predicate.OdometerRecord) *OdometerRecordDelete {
	ord.mutation.Where(ps...)
	return ord
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ord *OdometerRecordDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ord.sqlExec, ord.mutation, ord.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ord *OdometerRecordDelete) ExecX(ctx context.Context) int {
	n, err := ord.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ord *OdometerRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(odometerrecord.Table, sqlgraph.NewFieldSpec(odometerrecord.FieldID, field.TypeUUID))
	if ps := ord.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ord.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ord.mutation.done = true
	return affected, err
}

// OdometerRecordDeleteOne is the builder for deleting a single OdometerRecord entity.
type OdometerRecordDeleteOne struct {
	ord *OdometerRecordDelete
}

// Where appends a list predicates to the OdometerRecordDelete builder.
func (ordo *OdometerRecordDeleteOne) Where(ps ...predicate.OdometerRecord) *OdometerRecordDeleteOne {
	ordo.ord.mutation.Where(ps...)
	return ordo
}

// Exec executes the deletion query.
func (ordo *OdometerRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := ordo.ord.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{odometerrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ordo *OdometerRecordDeleteOne) ExecX(ctx context.Context) {
	if err := ordo.Exec(ctx); err != nil {
		panic(err)
	}
}
