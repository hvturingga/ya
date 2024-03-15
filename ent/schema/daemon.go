package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Daemon holds the schema definition for the Daemon entity.
type Daemon struct {
	ent.Schema
}

// Fields of the Daemon.
func (Daemon) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.Bool("enable").
			Default(false),
	}
}

// Edges of the Daemon.
func (Daemon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique(),
	}
}
