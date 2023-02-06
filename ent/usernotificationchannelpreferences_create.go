// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wallit/ent/user"
	"wallit/ent/usernotificationchannelpreferences"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserNotificationChannelPreferencesCreate is the builder for creating a UserNotificationChannelPreferences entity.
type UserNotificationChannelPreferencesCreate struct {
	config
	mutation *UserNotificationChannelPreferencesMutation
	hooks    []Hook
}

// SetChanel sets the "chanel" field.
func (uncpc *UserNotificationChannelPreferencesCreate) SetChanel(u usernotificationchannelpreferences.Chanel) *UserNotificationChannelPreferencesCreate {
	uncpc.mutation.SetChanel(u)
	return uncpc
}

// SetChanelUsersID sets the "chanel_users" edge to the User entity by ID.
func (uncpc *UserNotificationChannelPreferencesCreate) SetChanelUsersID(id int) *UserNotificationChannelPreferencesCreate {
	uncpc.mutation.SetChanelUsersID(id)
	return uncpc
}

// SetNillableChanelUsersID sets the "chanel_users" edge to the User entity by ID if the given value is not nil.
func (uncpc *UserNotificationChannelPreferencesCreate) SetNillableChanelUsersID(id *int) *UserNotificationChannelPreferencesCreate {
	if id != nil {
		uncpc = uncpc.SetChanelUsersID(*id)
	}
	return uncpc
}

// SetChanelUsers sets the "chanel_users" edge to the User entity.
func (uncpc *UserNotificationChannelPreferencesCreate) SetChanelUsers(u *User) *UserNotificationChannelPreferencesCreate {
	return uncpc.SetChanelUsersID(u.ID)
}

// Mutation returns the UserNotificationChannelPreferencesMutation object of the builder.
func (uncpc *UserNotificationChannelPreferencesCreate) Mutation() *UserNotificationChannelPreferencesMutation {
	return uncpc.mutation
}

// Save creates the UserNotificationChannelPreferences in the database.
func (uncpc *UserNotificationChannelPreferencesCreate) Save(ctx context.Context) (*UserNotificationChannelPreferences, error) {
	return withHooks[*UserNotificationChannelPreferences, UserNotificationChannelPreferencesMutation](ctx, uncpc.sqlSave, uncpc.mutation, uncpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uncpc *UserNotificationChannelPreferencesCreate) SaveX(ctx context.Context) *UserNotificationChannelPreferences {
	v, err := uncpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uncpc *UserNotificationChannelPreferencesCreate) Exec(ctx context.Context) error {
	_, err := uncpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uncpc *UserNotificationChannelPreferencesCreate) ExecX(ctx context.Context) {
	if err := uncpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uncpc *UserNotificationChannelPreferencesCreate) check() error {
	if _, ok := uncpc.mutation.Chanel(); !ok {
		return &ValidationError{Name: "chanel", err: errors.New(`ent: missing required field "UserNotificationChannelPreferences.chanel"`)}
	}
	if v, ok := uncpc.mutation.Chanel(); ok {
		if err := usernotificationchannelpreferences.ChanelValidator(v); err != nil {
			return &ValidationError{Name: "chanel", err: fmt.Errorf(`ent: validator failed for field "UserNotificationChannelPreferences.chanel": %w`, err)}
		}
	}
	return nil
}

func (uncpc *UserNotificationChannelPreferencesCreate) sqlSave(ctx context.Context) (*UserNotificationChannelPreferences, error) {
	if err := uncpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uncpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uncpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uncpc.mutation.id = &_node.ID
	uncpc.mutation.done = true
	return _node, nil
}

func (uncpc *UserNotificationChannelPreferencesCreate) createSpec() (*UserNotificationChannelPreferences, *sqlgraph.CreateSpec) {
	var (
		_node = &UserNotificationChannelPreferences{config: uncpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usernotificationchannelpreferences.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernotificationchannelpreferences.FieldID,
			},
		}
	)
	if value, ok := uncpc.mutation.Chanel(); ok {
		_spec.SetField(usernotificationchannelpreferences.FieldChanel, field.TypeEnum, value)
		_node.Chanel = value
	}
	if nodes := uncpc.mutation.ChanelUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usernotificationchannelpreferences.ChanelUsersTable,
			Columns: []string{usernotificationchannelpreferences.ChanelUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_notification_channels = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserNotificationChannelPreferencesCreateBulk is the builder for creating many UserNotificationChannelPreferences entities in bulk.
type UserNotificationChannelPreferencesCreateBulk struct {
	config
	builders []*UserNotificationChannelPreferencesCreate
}

// Save creates the UserNotificationChannelPreferences entities in the database.
func (uncpcb *UserNotificationChannelPreferencesCreateBulk) Save(ctx context.Context) ([]*UserNotificationChannelPreferences, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uncpcb.builders))
	nodes := make([]*UserNotificationChannelPreferences, len(uncpcb.builders))
	mutators := make([]Mutator, len(uncpcb.builders))
	for i := range uncpcb.builders {
		func(i int, root context.Context) {
			builder := uncpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserNotificationChannelPreferencesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uncpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uncpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, uncpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uncpcb *UserNotificationChannelPreferencesCreateBulk) SaveX(ctx context.Context) []*UserNotificationChannelPreferences {
	v, err := uncpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uncpcb *UserNotificationChannelPreferencesCreateBulk) Exec(ctx context.Context) error {
	_, err := uncpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uncpcb *UserNotificationChannelPreferencesCreateBulk) ExecX(ctx context.Context) {
	if err := uncpcb.Exec(ctx); err != nil {
		panic(err)
	}
}