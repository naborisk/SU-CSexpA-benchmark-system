// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/taskresult"
)

// TaskResultDelete is the builder for deleting a TaskResult entity.
type TaskResultDelete struct {
	config
	hooks    []Hook
	mutation *TaskResultMutation
}

// Where appends a list predicates to the TaskResultDelete builder.
func (trd *TaskResultDelete) Where(ps ...predicate.TaskResult) *TaskResultDelete {
	trd.mutation.Where(ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TaskResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, trd.sqlExec, trd.mutation, trd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TaskResultDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TaskResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(taskresult.Table, sqlgraph.NewFieldSpec(taskresult.FieldID, field.TypeInt))
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	trd.mutation.done = true
	return affected, err
}

// TaskResultDeleteOne is the builder for deleting a single TaskResult entity.
type TaskResultDeleteOne struct {
	trd *TaskResultDelete
}

// Where appends a list predicates to the TaskResultDelete builder.
func (trdo *TaskResultDeleteOne) Where(ps ...predicate.TaskResult) *TaskResultDeleteOne {
	trdo.trd.mutation.Where(ps...)
	return trdo
}

// Exec executes the deletion query.
func (trdo *TaskResultDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TaskResultDeleteOne) ExecX(ctx context.Context) {
	if err := trdo.Exec(ctx); err != nil {
		panic(err)
	}
}
