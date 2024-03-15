// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hvturingga/ya/ent/provider"
	"github.com/hvturingga/ya/ent/subscribe"
	"github.com/hvturingga/ya/ent/user"
)

// Subscribe is the model entity for the Subscribe schema.
type Subscribe struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Conf holds the value of the "conf" field.
	Conf string `json:"conf,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubscribeQuery when eager-loading is set.
	Edges               SubscribeEdges `json:"edges"`
	provider_subscribes *int
	selectValues        sql.SelectValues
}

// SubscribeEdges holds the relations/edges for other nodes in the graph.
type SubscribeEdges struct {
	// Provider holds the value of the provider edge.
	Provider *Provider `json:"provider,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"nodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProviderOrErr returns the Provider value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubscribeEdges) ProviderOrErr() (*Provider, error) {
	if e.loadedTypes[0] {
		if e.Provider == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: provider.Label}
		}
		return e.Provider, nil
	}
	return nil, &NotLoadedError{edge: "provider"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubscribeEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e SubscribeEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[2] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscribe) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscribe.FieldID:
			values[i] = new(sql.NullInt64)
		case subscribe.FieldName, subscribe.FieldLink, subscribe.FieldConf:
			values[i] = new(sql.NullString)
		case subscribe.ForeignKeys[0]: // provider_subscribes
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscribe fields.
func (s *Subscribe) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscribe.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subscribe.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subscribe.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				s.Link = value.String
			}
		case subscribe.FieldConf:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field conf", values[i])
			} else if value.Valid {
				s.Conf = value.String
			}
		case subscribe.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field provider_subscribes", value)
			} else if value.Valid {
				s.provider_subscribes = new(int)
				*s.provider_subscribes = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subscribe.
// This includes values selected through modifiers, order, etc.
func (s *Subscribe) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryProvider queries the "provider" edge of the Subscribe entity.
func (s *Subscribe) QueryProvider() *ProviderQuery {
	return NewSubscribeClient(s.config).QueryProvider(s)
}

// QueryUser queries the "user" edge of the Subscribe entity.
func (s *Subscribe) QueryUser() *UserQuery {
	return NewSubscribeClient(s.config).QueryUser(s)
}

// QueryNodes queries the "nodes" edge of the Subscribe entity.
func (s *Subscribe) QueryNodes() *NodeQuery {
	return NewSubscribeClient(s.config).QueryNodes(s)
}

// Update returns a builder for updating this Subscribe.
// Note that you need to call Subscribe.Unwrap() before calling this method if this Subscribe
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscribe) Update() *SubscribeUpdateOne {
	return NewSubscribeClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subscribe entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscribe) Unwrap() *Subscribe {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subscribe is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscribe) String() string {
	var builder strings.Builder
	builder.WriteString("Subscribe(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(s.Link)
	builder.WriteString(", ")
	builder.WriteString("conf=")
	builder.WriteString(s.Conf)
	builder.WriteByte(')')
	return builder.String()
}

// Subscribes is a parsable slice of Subscribe.
type Subscribes []*Subscribe
