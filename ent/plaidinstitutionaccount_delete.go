// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wallit/ent/plaidinstitutionaccount"
	"wallit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaidInstitutionAccountDelete is the builder for deleting a PlaidInstitutionAccount entity.
type PlaidInstitutionAccountDelete struct {
	config
	hooks    []Hook
	mutation *PlaidInstitutionAccountMutation
}

// Where appends a list predicates to the PlaidInstitutionAccountDelete builder.
func (piad *PlaidInstitutionAccountDelete) Where(ps ...predicate.PlaidInstitutionAccount) *PlaidInstitutionAccountDelete {
	piad.mutation.Where(ps...)
	return piad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (piad *PlaidInstitutionAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PlaidInstitutionAccountMutation](ctx, piad.sqlExec, piad.mutation, piad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (piad *PlaidInstitutionAccountDelete) ExecX(ctx context.Context) int {
	n, err := piad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (piad *PlaidInstitutionAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: plaidinstitutionaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: plaidinstitutionaccount.FieldID,
			},
		},
	}
	if ps := piad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, piad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	piad.mutation.done = true
	return affected, err
}

// PlaidInstitutionAccountDeleteOne is the builder for deleting a single PlaidInstitutionAccount entity.
type PlaidInstitutionAccountDeleteOne struct {
	piad *PlaidInstitutionAccountDelete
}

// Where appends a list predicates to the PlaidInstitutionAccountDelete builder.
func (piado *PlaidInstitutionAccountDeleteOne) Where(ps ...predicate.PlaidInstitutionAccount) *PlaidInstitutionAccountDeleteOne {
	piado.piad.mutation.Where(ps...)
	return piado
}

// Exec executes the deletion query.
func (piado *PlaidInstitutionAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := piado.piad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{plaidinstitutionaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (piado *PlaidInstitutionAccountDeleteOne) ExecX(ctx context.Context) {
	if err := piado.Exec(ctx); err != nil {
		panic(err)
	}
}