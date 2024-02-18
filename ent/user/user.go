// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// EdgeSubscribe holds the string denoting the subscribe edge name in mutations.
	EdgeSubscribe = "subscribe"
	// EdgeDaemon holds the string denoting the daemon edge name in mutations.
	EdgeDaemon = "daemon"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "users"
	// ProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	ProviderInverseTable = "providers"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "provider_user"
	// SubscribeTable is the table that holds the subscribe relation/edge.
	SubscribeTable = "users"
	// SubscribeInverseTable is the table name for the Subscribe entity.
	// It exists in this package in order to avoid circular dependency with the "subscribe" package.
	SubscribeInverseTable = "subscribes"
	// SubscribeColumn is the table column denoting the subscribe relation/edge.
	SubscribeColumn = "subscribe_user"
	// DaemonTable is the table that holds the daemon relation/edge.
	DaemonTable = "users"
	// DaemonInverseTable is the table name for the Daemon entity.
	// It exists in this package in order to avoid circular dependency with the "daemon" package.
	DaemonInverseTable = "daemons"
	// DaemonColumn is the table column denoting the daemon relation/edge.
	DaemonColumn = "daemon_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldActive,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"daemon_user",
	"provider_user",
	"subscribe_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByProviderField orders the results by provider field.
func ByProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderStep(), sql.OrderByField(field, opts...))
	}
}

// BySubscribeField orders the results by subscribe field.
func BySubscribeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscribeStep(), sql.OrderByField(field, opts...))
	}
}

// ByDaemonField orders the results by daemon field.
func ByDaemonField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDaemonStep(), sql.OrderByField(field, opts...))
	}
}
func newProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ProviderTable, ProviderColumn),
	)
}
func newSubscribeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscribeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SubscribeTable, SubscribeColumn),
	)
}
func newDaemonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DaemonInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, DaemonTable, DaemonColumn),
	)
}
