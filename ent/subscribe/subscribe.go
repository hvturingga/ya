// Code generated by ent, DO NOT EDIT.

package subscribe

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subscribe type in the database.
	Label = "subscribe"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldConf holds the string denoting the conf field in the database.
	FieldConf = "conf"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeNodes holds the string denoting the nodes edge name in mutations.
	EdgeNodes = "nodes"
	// Table holds the table name of the subscribe in the database.
	Table = "subscribes"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "subscribes"
	// ProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	ProviderInverseTable = "providers"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "provider_subscribes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "subscribe_user"
	// NodesTable is the table that holds the nodes relation/edge.
	NodesTable = "nodes"
	// NodesInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodesInverseTable = "nodes"
	// NodesColumn is the table column denoting the nodes relation/edge.
	NodesColumn = "subscribe_nodes"
)

// Columns holds all SQL columns for subscribe fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLink,
	FieldConf,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subscribes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"provider_subscribes",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ConfValidator is a validator for the "conf" field. It is called by the builders before save.
	ConfValidator func(string) error
)

// OrderOption defines the ordering options for the Subscribe queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLink orders the results by the link field.
func ByLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLink, opts...).ToFunc()
}

// ByConf orders the results by the conf field.
func ByConf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConf, opts...).ToFunc()
}

// ByProviderField orders the results by provider field.
func ByProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByNodesCount orders the results by nodes count.
func ByNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodesStep(), opts...)
	}
}

// ByNodes orders the results by nodes terms.
func ByNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, UserTable, UserColumn),
	)
}
func newNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NodesTable, NodesColumn),
	)
}
