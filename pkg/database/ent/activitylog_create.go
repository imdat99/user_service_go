// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/pkg/database/ent/activitylog"
	"app/pkg/database/ent/user"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActivityLogCreate is the builder for creating a ActivityLog entity.
type ActivityLogCreate struct {
	config
	mutation *ActivityLogMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (alc *ActivityLogCreate) SetUserID(s string) *ActivityLogCreate {
	alc.mutation.SetUserID(s)
	return alc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (alc *ActivityLogCreate) SetNillableUserID(s *string) *ActivityLogCreate {
	if s != nil {
		alc.SetUserID(*s)
	}
	return alc
}

// SetActivityType sets the "activity_type" field.
func (alc *ActivityLogCreate) SetActivityType(at activitylog.ActivityType) *ActivityLogCreate {
	alc.mutation.SetActivityType(at)
	return alc
}

// SetDescription sets the "description" field.
func (alc *ActivityLogCreate) SetDescription(s string) *ActivityLogCreate {
	alc.mutation.SetDescription(s)
	return alc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (alc *ActivityLogCreate) SetNillableDescription(s *string) *ActivityLogCreate {
	if s != nil {
		alc.SetDescription(*s)
	}
	return alc
}

// SetIPAddress sets the "ip_address" field.
func (alc *ActivityLogCreate) SetIPAddress(s string) *ActivityLogCreate {
	alc.mutation.SetIPAddress(s)
	return alc
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (alc *ActivityLogCreate) SetNillableIPAddress(s *string) *ActivityLogCreate {
	if s != nil {
		alc.SetIPAddress(*s)
	}
	return alc
}

// SetUserAgent sets the "user_agent" field.
func (alc *ActivityLogCreate) SetUserAgent(s string) *ActivityLogCreate {
	alc.mutation.SetUserAgent(s)
	return alc
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (alc *ActivityLogCreate) SetNillableUserAgent(s *string) *ActivityLogCreate {
	if s != nil {
		alc.SetUserAgent(*s)
	}
	return alc
}

// SetMetadata sets the "metadata" field.
func (alc *ActivityLogCreate) SetMetadata(jm json.RawMessage) *ActivityLogCreate {
	alc.mutation.SetMetadata(jm)
	return alc
}

// SetCreatedAt sets the "created_at" field.
func (alc *ActivityLogCreate) SetCreatedAt(t time.Time) *ActivityLogCreate {
	alc.mutation.SetCreatedAt(t)
	return alc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alc *ActivityLogCreate) SetNillableCreatedAt(t *time.Time) *ActivityLogCreate {
	if t != nil {
		alc.SetCreatedAt(*t)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *ActivityLogCreate) SetID(s string) *ActivityLogCreate {
	alc.mutation.SetID(s)
	return alc
}

// SetUser sets the "user" edge to the User entity.
func (alc *ActivityLogCreate) SetUser(u *User) *ActivityLogCreate {
	return alc.SetUserID(u.ID)
}

// Mutation returns the ActivityLogMutation object of the builder.
func (alc *ActivityLogCreate) Mutation() *ActivityLogMutation {
	return alc.mutation
}

// Save creates the ActivityLog in the database.
func (alc *ActivityLogCreate) Save(ctx context.Context) (*ActivityLog, error) {
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *ActivityLogCreate) SaveX(ctx context.Context) *ActivityLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *ActivityLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *ActivityLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *ActivityLogCreate) check() error {
	if _, ok := alc.mutation.ActivityType(); !ok {
		return &ValidationError{Name: "activity_type", err: errors.New(`ent: missing required field "ActivityLog.activity_type"`)}
	}
	if v, ok := alc.mutation.ActivityType(); ok {
		if err := activitylog.ActivityTypeValidator(v); err != nil {
			return &ValidationError{Name: "activity_type", err: fmt.Errorf(`ent: validator failed for field "ActivityLog.activity_type": %w`, err)}
		}
	}
	return nil
}

func (alc *ActivityLogCreate) sqlSave(ctx context.Context) (*ActivityLog, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ActivityLog.ID type: %T", _spec.ID.Value)
		}
	}
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *ActivityLogCreate) createSpec() (*ActivityLog, *sqlgraph.CreateSpec) {
	var (
		_node = &ActivityLog{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(activitylog.Table, sqlgraph.NewFieldSpec(activitylog.FieldID, field.TypeString))
	)
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := alc.mutation.ActivityType(); ok {
		_spec.SetField(activitylog.FieldActivityType, field.TypeEnum, value)
		_node.ActivityType = value
	}
	if value, ok := alc.mutation.Description(); ok {
		_spec.SetField(activitylog.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := alc.mutation.IPAddress(); ok {
		_spec.SetField(activitylog.FieldIPAddress, field.TypeString, value)
		_node.IPAddress = value
	}
	if value, ok := alc.mutation.UserAgent(); ok {
		_spec.SetField(activitylog.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = value
	}
	if value, ok := alc.mutation.Metadata(); ok {
		_spec.SetField(activitylog.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.SetField(activitylog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := alc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activitylog.UserTable,
			Columns: []string{activitylog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActivityLogCreateBulk is the builder for creating many ActivityLog entities in bulk.
type ActivityLogCreateBulk struct {
	config
	err      error
	builders []*ActivityLogCreate
}

// Save creates the ActivityLog entities in the database.
func (alcb *ActivityLogCreateBulk) Save(ctx context.Context) ([]*ActivityLog, error) {
	if alcb.err != nil {
		return nil, alcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*ActivityLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActivityLogMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *ActivityLogCreateBulk) SaveX(ctx context.Context) []*ActivityLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *ActivityLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *ActivityLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}
