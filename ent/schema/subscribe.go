package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subscribe holds the schema definition for the Subscribe entity.
type Subscribe struct {
	ent.Schema
}

// Fields of the Subscribe.
func (Subscribe) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			MaxLen(24).
			NotEmpty(),
		field.
			String("link").
			Optional(),
		field.
			String("conf").
			NotEmpty(),
	}
}

// Edges of the Subscribe.
func (Subscribe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", Provider.Type).
			Ref("subscribes").
			Unique(),
		edge.To("user", User.Type).
			Unique(),
		edge.To("nodes", Node.Type),
	}
}
