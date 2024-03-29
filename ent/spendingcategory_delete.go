// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wallit/ent/predicate"
	"wallit/ent/spendingcategory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpendingCategoryDelete is the builder for deleting a SpendingCategory entity.
type SpendingCategoryDelete struct {
	config
	hooks    []Hook
	mutation *SpendingCategoryMutation
}

// Where appends a list predicates to the SpendingCategoryDelete builder.
func (scd *SpendingCategoryDelete) Where(ps ...predicate.SpendingCategory) *SpendingCategoryDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SpendingCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SpendingCategoryMutation](ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SpendingCategoryDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SpendingCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: spendingcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spendingcategory.FieldID,
			},
		},
	}
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	scd.mutation.done = true
	return affected, err
}

// SpendingCategoryDeleteOne is the builder for deleting a single SpendingCategory entity.
type SpendingCategoryDeleteOne struct {
	scd *SpendingCategoryDelete
}

// Where appends a list predicates to the SpendingCategoryDelete builder.
func (scdo *SpendingCategoryDeleteOne) Where(ps ...predicate.SpendingCategory) *SpendingCategoryDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *SpendingCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spendingcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SpendingCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}
