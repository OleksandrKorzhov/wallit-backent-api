// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wallit/ent/predicate"
	"wallit/ent/usernotificationchannelpreferences"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserNotificationChannelPreferencesDelete is the builder for deleting a UserNotificationChannelPreferences entity.
type UserNotificationChannelPreferencesDelete struct {
	config
	hooks    []Hook
	mutation *UserNotificationChannelPreferencesMutation
}

// Where appends a list predicates to the UserNotificationChannelPreferencesDelete builder.
func (uncpd *UserNotificationChannelPreferencesDelete) Where(ps ...predicate.UserNotificationChannelPreferences) *UserNotificationChannelPreferencesDelete {
	uncpd.mutation.Where(ps...)
	return uncpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uncpd *UserNotificationChannelPreferencesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserNotificationChannelPreferencesMutation](ctx, uncpd.sqlExec, uncpd.mutation, uncpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uncpd *UserNotificationChannelPreferencesDelete) ExecX(ctx context.Context) int {
	n, err := uncpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uncpd *UserNotificationChannelPreferencesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usernotificationchannelpreferences.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernotificationchannelpreferences.FieldID,
			},
		},
	}
	if ps := uncpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uncpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uncpd.mutation.done = true
	return affected, err
}

// UserNotificationChannelPreferencesDeleteOne is the builder for deleting a single UserNotificationChannelPreferences entity.
type UserNotificationChannelPreferencesDeleteOne struct {
	uncpd *UserNotificationChannelPreferencesDelete
}

// Where appends a list predicates to the UserNotificationChannelPreferencesDelete builder.
func (uncpdo *UserNotificationChannelPreferencesDeleteOne) Where(ps ...predicate.UserNotificationChannelPreferences) *UserNotificationChannelPreferencesDeleteOne {
	uncpdo.uncpd.mutation.Where(ps...)
	return uncpdo
}

// Exec executes the deletion query.
func (uncpdo *UserNotificationChannelPreferencesDeleteOne) Exec(ctx context.Context) error {
	n, err := uncpdo.uncpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usernotificationchannelpreferences.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uncpdo *UserNotificationChannelPreferencesDeleteOne) ExecX(ctx context.Context) {
	if err := uncpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
