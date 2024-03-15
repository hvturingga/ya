// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hvturingga/ya/ent/daemon"
	"github.com/hvturingga/ya/ent/provider"
	"github.com/hvturingga/ya/ent/subscribe"
	"github.com/hvturingga/ya/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges          UserEdges `json:"edges"`
	daemon_user    *int
	provider_user  *int
	subscribe_user *int
	selectValues   sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Provider holds the value of the provider edge.
	Provider *Provider `json:"provider,omitempty"`
	// Subscribe holds the value of the subscribe edge.
	Subscribe *Subscribe `json:"subscribe,omitempty"`
	// Daemon holds the value of the daemon edge.
	Daemon *Daemon `json:"daemon,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProviderOrErr returns the Provider value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProviderOrErr() (*Provider, error) {
	if e.loadedTypes[0] {
		if e.Provider == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: provider.Label}
		}
		return e.Provider, nil
	}
	return nil, &NotLoadedError{edge: "provider"}
}

// SubscribeOrErr returns the Subscribe value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SubscribeOrErr() (*Subscribe, error) {
	if e.loadedTypes[1] {
		if e.Subscribe == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subscribe.Label}
		}
		return e.Subscribe, nil
	}
	return nil, &NotLoadedError{edge: "subscribe"}
}

// DaemonOrErr returns the Daemon value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) DaemonOrErr() (*Daemon, error) {
	if e.loadedTypes[2] {
		if e.Daemon == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: daemon.Label}
		}
		return e.Daemon, nil
	}
	return nil, &NotLoadedError{edge: "daemon"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[0]: // daemon_user
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[1]: // provider_user
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[2]: // subscribe_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field daemon_user", value)
			} else if value.Valid {
				u.daemon_user = new(int)
				*u.daemon_user = int(value.Int64)
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field provider_user", value)
			} else if value.Valid {
				u.provider_user = new(int)
				*u.provider_user = int(value.Int64)
			}
		case user.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field subscribe_user", value)
			} else if value.Valid {
				u.subscribe_user = new(int)
				*u.subscribe_user = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryProvider queries the "provider" edge of the User entity.
func (u *User) QueryProvider() *ProviderQuery {
	return NewUserClient(u.config).QueryProvider(u)
}

// QuerySubscribe queries the "subscribe" edge of the User entity.
func (u *User) QuerySubscribe() *SubscribeQuery {
	return NewUserClient(u.config).QuerySubscribe(u)
}

// QueryDaemon queries the "daemon" edge of the User entity.
func (u *User) QueryDaemon() *DaemonQuery {
	return NewUserClient(u.config).QueryDaemon(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
