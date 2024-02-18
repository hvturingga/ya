package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
		field.
			Bool("active").
			Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", Provider.Type).
			Ref("user").
			Unique(),
		edge.From("subscribe", Subscribe.Type).
			Ref("user").
			Unique(),
		edge.From("daemon", Daemon.Type).
			Ref("user").
			Unique(),
	}
}
