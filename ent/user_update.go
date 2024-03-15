// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hvturingga/ya/ent/daemon"
	"github.com/hvturingga/ya/ent/predicate"
	"github.com/hvturingga/ya/ent/provider"
	"github.com/hvturingga/ya/ent/subscribe"
	"github.com/hvturingga/ya/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (uu *UserUpdate) SetProviderID(id int) *UserUpdate {
	uu.mutation.SetProviderID(id)
	return uu
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableProviderID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetProviderID(*id)
	}
	return uu
}

// SetProvider sets the "provider" edge to the Provider entity.
func (uu *UserUpdate) SetProvider(p *Provider) *UserUpdate {
	return uu.SetProviderID(p.ID)
}

// SetSubscribeID sets the "subscribe" edge to the Subscribe entity by ID.
func (uu *UserUpdate) SetSubscribeID(id int) *UserUpdate {
	uu.mutation.SetSubscribeID(id)
	return uu
}

// SetNillableSubscribeID sets the "subscribe" edge to the Subscribe entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableSubscribeID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetSubscribeID(*id)
	}
	return uu
}

// SetSubscribe sets the "subscribe" edge to the Subscribe entity.
func (uu *UserUpdate) SetSubscribe(s *Subscribe) *UserUpdate {
	return uu.SetSubscribeID(s.ID)
}

// SetDaemonID sets the "daemon" edge to the Daemon entity by ID.
func (uu *UserUpdate) SetDaemonID(id int) *UserUpdate {
	uu.mutation.SetDaemonID(id)
	return uu
}

// SetNillableDaemonID sets the "daemon" edge to the Daemon entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableDaemonID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetDaemonID(*id)
	}
	return uu
}

// SetDaemon sets the "daemon" edge to the Daemon entity.
func (uu *UserUpdate) SetDaemon(d *Daemon) *UserUpdate {
	return uu.SetDaemonID(d.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (uu *UserUpdate) ClearProvider() *UserUpdate {
	uu.mutation.ClearProvider()
	return uu
}

// ClearSubscribe clears the "subscribe" edge to the Subscribe entity.
func (uu *UserUpdate) ClearSubscribe() *UserUpdate {
	uu.mutation.ClearSubscribe()
	return uu
}

// ClearDaemon clears the "daemon" edge to the Daemon entity.
func (uu *UserUpdate) ClearDaemon() *UserUpdate {
	uu.mutation.ClearDaemon()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uu.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.ProviderTable,
			Columns: []string{user.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.ProviderTable,
			Columns: []string{user.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.SubscribeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.SubscribeTable,
			Columns: []string{user.SubscribeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscribe.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SubscribeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.SubscribeTable,
			Columns: []string{user.SubscribeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscribe.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.DaemonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.DaemonTable,
			Columns: []string{user.DaemonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(daemon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DaemonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.DaemonTable,
			Columns: []string{user.DaemonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(daemon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (uuo *UserUpdateOne) SetProviderID(id int) *UserUpdateOne {
	uuo.mutation.SetProviderID(id)
	return uuo
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProviderID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetProviderID(*id)
	}
	return uuo
}

// SetProvider sets the "provider" edge to the Provider entity.
func (uuo *UserUpdateOne) SetProvider(p *Provider) *UserUpdateOne {
	return uuo.SetProviderID(p.ID)
}

// SetSubscribeID sets the "subscribe" edge to the Subscribe entity by ID.
func (uuo *UserUpdateOne) SetSubscribeID(id int) *UserUpdateOne {
	uuo.mutation.SetSubscribeID(id)
	return uuo
}

// SetNillableSubscribeID sets the "subscribe" edge to the Subscribe entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSubscribeID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSubscribeID(*id)
	}
	return uuo
}

// SetSubscribe sets the "subscribe" edge to the Subscribe entity.
func (uuo *UserUpdateOne) SetSubscribe(s *Subscribe) *UserUpdateOne {
	return uuo.SetSubscribeID(s.ID)
}

// SetDaemonID sets the "daemon" edge to the Daemon entity by ID.
func (uuo *UserUpdateOne) SetDaemonID(id int) *UserUpdateOne {
	uuo.mutation.SetDaemonID(id)
	return uuo
}

// SetNillableDaemonID sets the "daemon" edge to the Daemon entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDaemonID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetDaemonID(*id)
	}
	return uuo
}

// SetDaemon sets the "daemon" edge to the Daemon entity.
func (uuo *UserUpdateOne) SetDaemon(d *Daemon) *UserUpdateOne {
	return uuo.SetDaemonID(d.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (uuo *UserUpdateOne) ClearProvider() *UserUpdateOne {
	uuo.mutation.ClearProvider()
	return uuo
}

// ClearSubscribe clears the "subscribe" edge to the Subscribe entity.
func (uuo *UserUpdateOne) ClearSubscribe() *UserUpdateOne {
	uuo.mutation.ClearSubscribe()
	return uuo
}

// ClearDaemon clears the "daemon" edge to the Daemon entity.
func (uuo *UserUpdateOne) ClearDaemon() *UserUpdateOne {
	uuo.mutation.ClearDaemon()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uuo.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.ProviderTable,
			Columns: []string{user.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.ProviderTable,
			Columns: []string{user.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.SubscribeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.SubscribeTable,
			Columns: []string{user.SubscribeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscribe.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SubscribeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.SubscribeTable,
			Columns: []string{user.SubscribeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscribe.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.DaemonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.DaemonTable,
			Columns: []string{user.DaemonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(daemon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DaemonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.DaemonTable,
			Columns: []string{user.DaemonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(daemon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
