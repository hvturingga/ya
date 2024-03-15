package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Provider holds the schema definition for the Provider entity.
type Provider struct {
	ent.Schema
}

// Fields of the Provider.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			NotEmpty(),
		field.
			String("version").
			Optional(),
		field.
			String("path").
			Optional(),
	}
}

// Edges of the Provider.
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscribes", Subscribe.Type),
		edge.To("user", User.Type).
			Unique(),
	}
}
