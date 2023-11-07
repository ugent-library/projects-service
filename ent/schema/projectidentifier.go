package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ProjectIdentifier struct {
	ent.Schema
}

func (ProjectIdentifier) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").
			Unique(),
	}
}

func (ProjectIdentifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
	}
}
