// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/pkg/database/ent/activitylog"
	"app/pkg/database/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActivityLogDelete is the builder for deleting a ActivityLog entity.
type ActivityLogDelete struct {
	config
	hooks    []Hook
	mutation *ActivityLogMutation
}

// Where appends a list predicates to the ActivityLogDelete builder.
func (ald *ActivityLogDelete) Where(ps ...predicate.ActivityLog) *ActivityLogDelete {
	ald.mutation.Where(ps...)
	return ald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ald *ActivityLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ald.sqlExec, ald.mutation, ald.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ald *ActivityLogDelete) ExecX(ctx context.Context) int {
	n, err := ald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ald *ActivityLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(activitylog.Table, sqlgraph.NewFieldSpec(activitylog.FieldID, field.TypeString))
	if ps := ald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ald.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ald.mutation.done = true
	return affected, err
}

// ActivityLogDeleteOne is the builder for deleting a single ActivityLog entity.
type ActivityLogDeleteOne struct {
	ald *ActivityLogDelete
}

// Where appends a list predicates to the ActivityLogDelete builder.
func (aldo *ActivityLogDeleteOne) Where(ps ...predicate.ActivityLog) *ActivityLogDeleteOne {
	aldo.ald.mutation.Where(ps...)
	return aldo
}

// Exec executes the deletion query.
func (aldo *ActivityLogDeleteOne) Exec(ctx context.Context) error {
	n, err := aldo.ald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{activitylog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aldo *ActivityLogDeleteOne) ExecX(ctx context.Context) {
	if err := aldo.Exec(ctx); err != nil {
		panic(err)
	}
}
