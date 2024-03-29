// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hvturingga/ya/ent/node"
	"github.com/hvturingga/ya/ent/predicate"
	"github.com/hvturingga/ya/ent/provider"
	"github.com/hvturingga/ya/ent/subscribe"
	"github.com/hvturingga/ya/ent/user"
)

// SubscribeUpdate is the builder for updating Subscribe entities.
type SubscribeUpdate struct {
	config
	hooks    []Hook
	mutation *SubscribeMutation
}

// Where appends a list predicates to the SubscribeUpdate builder.
func (su *SubscribeUpdate) Where(ps ...predicate.Subscribe) *SubscribeUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SubscribeUpdate) SetName(s string) *SubscribeUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SubscribeUpdate) SetNillableName(s *string) *SubscribeUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetLink sets the "link" field.
func (su *SubscribeUpdate) SetLink(s string) *SubscribeUpdate {
	su.mutation.SetLink(s)
	return su
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (su *SubscribeUpdate) SetNillableLink(s *string) *SubscribeUpdate {
	if s != nil {
		su.SetLink(*s)
	}
	return su
}

// ClearLink clears the value of the "link" field.
func (su *SubscribeUpdate) ClearLink() *SubscribeUpdate {
	su.mutation.ClearLink()
	return su
}

// SetConf sets the "conf" field.
func (su *SubscribeUpdate) SetConf(s string) *SubscribeUpdate {
	su.mutation.SetConf(s)
	return su
}

// SetNillableConf sets the "conf" field if the given value is not nil.
func (su *SubscribeUpdate) SetNillableConf(s *string) *SubscribeUpdate {
	if s != nil {
		su.SetConf(*s)
	}
	return su
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (su *SubscribeUpdate) SetProviderID(id int) *SubscribeUpdate {
	su.mutation.SetProviderID(id)
	return su
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (su *SubscribeUpdate) SetNillableProviderID(id *int) *SubscribeUpdate {
	if id != nil {
		su = su.SetProviderID(*id)
	}
	return su
}

// SetProvider sets the "provider" edge to the Provider entity.
func (su *SubscribeUpdate) SetProvider(p *Provider) *SubscribeUpdate {
	return su.SetProviderID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SubscribeUpdate) SetUserID(id int) *SubscribeUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SubscribeUpdate) SetNillableUserID(id *int) *SubscribeUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SubscribeUpdate) SetUser(u *User) *SubscribeUpdate {
	return su.SetUserID(u.ID)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (su *SubscribeUpdate) AddNodeIDs(ids ...int) *SubscribeUpdate {
	su.mutation.AddNodeIDs(ids...)
	return su
}

// AddNodes adds the "nodes" edges to the Node entity.
func (su *SubscribeUpdate) AddNodes(n ...*Node) *SubscribeUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return su.AddNodeIDs(ids...)
}

// Mutation returns the SubscribeMutation object of the builder.
func (su *SubscribeUpdate) Mutation() *SubscribeMutation {
	return su.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (su *SubscribeUpdate) ClearProvider() *SubscribeUpdate {
	su.mutation.ClearProvider()
	return su
}

// ClearUser clears the "user" edge to the User entity.
func (su *SubscribeUpdate) ClearUser() *SubscribeUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (su *SubscribeUpdate) ClearNodes() *SubscribeUpdate {
	su.mutation.ClearNodes()
	return su
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (su *SubscribeUpdate) RemoveNodeIDs(ids ...int) *SubscribeUpdate {
	su.mutation.RemoveNodeIDs(ids...)
	return su
}

// RemoveNodes removes "nodes" edges to Node entities.
func (su *SubscribeUpdate) RemoveNodes(n ...*Node) *SubscribeUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return su.RemoveNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscribeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscribeUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscribeUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscribeUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubscribeUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := subscribe.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subscribe.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Conf(); ok {
		if err := subscribe.ConfValidator(v); err != nil {
			return &ValidationError{Name: "conf", err: fmt.Errorf(`ent: validator failed for field "Subscribe.conf": %w`, err)}
		}
	}
	return nil
}

func (su *SubscribeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(subscribe.Table, subscribe.Columns, sqlgraph.NewFieldSpec(subscribe.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(subscribe.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Link(); ok {
		_spec.SetField(subscribe.FieldLink, field.TypeString, value)
	}
	if su.mutation.LinkCleared() {
		_spec.ClearField(subscribe.FieldLink, field.TypeString)
	}
	if value, ok := su.mutation.Conf(); ok {
		_spec.SetField(subscribe.FieldConf, field.TypeString, value)
	}
	if su.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscribe.ProviderTable,
			Columns: []string{subscribe.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscribe.ProviderTable,
			Columns: []string{subscribe.ProviderColumn},
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
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   subscribe.UserTable,
			Columns: []string{subscribe.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   subscribe.UserTable,
			Columns: []string{subscribe.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscribe.NodesTable,
			Columns: []string{subscribe.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedNodesIDs(); len(nodes) > 0 && !su.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscribe.NodesTable,
			Columns: []string{subscribe.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscribe.NodesTable,
			Columns: []string{subscribe.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscribe.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubscribeUpdateOne is the builder for updating a single Subscribe entity.
type SubscribeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscribeMutation
}

// SetName sets the "name" field.
func (suo *SubscribeUpdateOne) SetName(s string) *SubscribeUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SubscribeUpdateOne) SetNillableName(s *string) *SubscribeUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetLink sets the "link" field.
func (suo *SubscribeUpdateOne) SetLink(s string) *SubscribeUpdateOne {
	suo.mutation.SetLink(s)
	return suo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (suo *SubscribeUpdateOne) SetNillableLink(s *string) *SubscribeUpdateOne {
	if s != nil {
		suo.SetLink(*s)
	}
	return suo
}

// ClearLink clears the value of the "link" field.
func (suo *SubscribeUpdateOne) ClearLink() *SubscribeUpdateOne {
	suo.mutation.ClearLink()
	return suo
}

// SetConf sets the "conf" field.
func (suo *SubscribeUpdateOne) SetConf(s string) *SubscribeUpdateOne {
	suo.mutation.SetConf(s)
	return suo
}

// SetNillableConf sets the "conf" field if the given value is not nil.
func (suo *SubscribeUpdateOne) SetNillableConf(s *string) *SubscribeUpdateOne {
	if s != nil {
		suo.SetConf(*s)
	}
	return suo
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (suo *SubscribeUpdateOne) SetProviderID(id int) *SubscribeUpdateOne {
	suo.mutation.SetProviderID(id)
	return suo
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (suo *SubscribeUpdateOne) SetNillableProviderID(id *int) *SubscribeUpdateOne {
	if id != nil {
		suo = suo.SetProviderID(*id)
	}
	return suo
}

// SetProvider sets the "provider" edge to the Provider entity.
func (suo *SubscribeUpdateOne) SetProvider(p *Provider) *SubscribeUpdateOne {
	return suo.SetProviderID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SubscribeUpdateOne) SetUserID(id int) *SubscribeUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SubscribeUpdateOne) SetNillableUserID(id *int) *SubscribeUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SubscribeUpdateOne) SetUser(u *User) *SubscribeUpdateOne {
	return suo.SetUserID(u.ID)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (suo *SubscribeUpdateOne) AddNodeIDs(ids ...int) *SubscribeUpdateOne {
	suo.mutation.AddNodeIDs(ids...)
	return suo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (suo *SubscribeUpdateOne) AddNodes(n ...*Node) *SubscribeUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return suo.AddNodeIDs(ids...)
}

// Mutation returns the SubscribeMutation object of the builder.
func (suo *SubscribeUpdateOne) Mutation() *SubscribeMutation {
	return suo.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (suo *SubscribeUpdateOne) ClearProvider() *SubscribeUpdateOne {
	suo.mutation.ClearProvider()
	return suo
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SubscribeUpdateOne) ClearUser() *SubscribeUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (suo *SubscribeUpdateOne) ClearNodes() *SubscribeUpdateOne {
	suo.mutation.ClearNodes()
	return suo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (suo *SubscribeUpdateOne) RemoveNodeIDs(ids ...int) *SubscribeUpdateOne {
	suo.mutation.RemoveNodeIDs(ids...)
	return suo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (suo *SubscribeUpdateOne) RemoveNodes(n ...*Node) *SubscribeUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return suo.RemoveNodeIDs(ids...)
}

// Where appends a list predicates to the SubscribeUpdate builder.
func (suo *SubscribeUpdateOne) Where(ps ...predicate.Subscribe) *SubscribeUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscribeUpdateOne) Select(field string, fields ...string) *SubscribeUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscribe entity.
func (suo *SubscribeUpdateOne) Save(ctx context.Context) (*Subscribe, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscribeUpdateOne) SaveX(ctx context.Context) *Subscribe {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscribeUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscribeUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubscribeUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := subscribe.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subscribe.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Conf(); ok {
		if err := subscribe.ConfValidator(v); err != nil {
			return &ValidationError{Name: "conf", err: fmt.Errorf(`ent: validator failed for field "Subscribe.conf": %w`, err)}
		}
	}
	return nil
}

func (suo *SubscribeUpdateOne) sqlSave(ctx context.Context) (_node *Subscribe, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(subscribe.Table, subscribe.Columns, sqlgraph.NewFieldSpec(subscribe.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscribe.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscribe.FieldID)
		for _, f := range fields {
			if !subscribe.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscribe.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(subscribe.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Link(); ok {
		_spec.SetField(subscribe.FieldLink, field.TypeString, value)
	}
	if suo.mutation.LinkCleared() {
		_spec.ClearField(subscribe.FieldLink, field.TypeString)
	}
	if value, ok := suo.mutation.Conf(); ok {
		_spec.SetField(subscribe.FieldConf, field.TypeString, value)
	}
	if suo.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscribe.ProviderTable,
			Columns: []string{subscribe.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscribe.ProviderTable,
			Columns: []string{subscribe.ProviderColumn},
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
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   subscribe.UserTable,
			Columns: []string{subscribe.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   subscribe.UserTable,
			Columns: []string{subscribe.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscribe.NodesTable,
			Columns: []string{subscribe.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !suo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscribe.NodesTable,
			Columns: []string{subscribe.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscribe.NodesTable,
			Columns: []string{subscribe.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subscribe{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscribe.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
