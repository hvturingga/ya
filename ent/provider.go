// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hvturingga/ya/ent/provider"
	"github.com/hvturingga/ya/ent/user"
)

// Provider is the model entity for the Provider schema.
type Provider struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderQuery when eager-loading is set.
	Edges        ProviderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProviderEdges holds the relations/edges for other nodes in the graph.
type ProviderEdges struct {
	// Subscribes holds the value of the subscribes edge.
	Subscribes []*Subscribe `json:"subscribes,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SubscribesOrErr returns the Subscribes value or an error if the edge
// was not loaded in eager-loading.
func (e ProviderEdges) SubscribesOrErr() ([]*Subscribe, error) {
	if e.loadedTypes[0] {
		return e.Subscribes, nil
	}
	return nil, &NotLoadedError{edge: "subscribes"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProviderEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case provider.FieldID:
			values[i] = new(sql.NullInt64)
		case provider.FieldName, provider.FieldVersion, provider.FieldPath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provider fields.
func (pr *Provider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case provider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case provider.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				pr.Version = value.String
			}
		case provider.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				pr.Path = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Provider.
// This includes values selected through modifiers, order, etc.
func (pr *Provider) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QuerySubscribes queries the "subscribes" edge of the Provider entity.
func (pr *Provider) QuerySubscribes() *SubscribeQuery {
	return NewProviderClient(pr.config).QuerySubscribes(pr)
}

// QueryUser queries the "user" edge of the Provider entity.
func (pr *Provider) QueryUser() *UserQuery {
	return NewProviderClient(pr.config).QueryUser(pr)
}

// Update returns a builder for updating this Provider.
// Note that you need to call Provider.Unwrap() before calling this method if this Provider
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provider) Update() *ProviderUpdateOne {
	return NewProviderClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Provider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provider) Unwrap() *Provider {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Provider is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provider) String() string {
	var builder strings.Builder
	builder.WriteString("Provider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(pr.Version)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(pr.Path)
	builder.WriteByte(')')
	return builder.String()
}

// Providers is a parsable slice of Provider.
type Providers []*Provider
