package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

var timeUTC = func() time.Time {
	return time.Now().UTC()
}

type Project struct {
	ent.Schema
}

type Identifier struct {
	Value map[string][]string `json:"value"`
}

type TranslatedString struct {
	Value map[string]string `json:"value"`
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Int("project_identifier_id"),
		field.JSON("identifier", Identifier{}).
			Default(Identifier{}),
		field.JSON("name", TranslatedString{}).
			Default(TranslatedString{}),
		field.JSON("description", TranslatedString{}).
			Default(TranslatedString{}),
		field.String("founding_date").
			Optional(),
		field.String("dissolution_date").
			Optional(),
		field.String("acronym").
			Optional(),
		field.String("grant_id").
			Optional(),
		field.String("funding_programme").
			Optional(),
		field.Bool("deleted").
			Default(false),
		field.Time("created_at").
			Default(timeUTC).
			Immutable(),
		field.Time("updated_at").
			Default(timeUTC).
			UpdateDefault(timeUTC),
		// field.String("ts").
		// 	Optional().
		// 	SchemaType(map[string]string{
		// 		dialect.Postgres: "tsvector NULL GENERATED ALWAYS AS ((to_tsvector('simple'::regconfig, jsonb_path_query_array(identifier, '$.**{2}'::jsonpath)) || to_tsvector('simple'::regconfig, (id)::text)) || to_tsvector('usimple'::regconfig, jsonb_path_query_array(name, '$.**{2}'::jsonpath))) STORED",
		// 	}),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("identifiedBy", ProjectIdentifier.Type).
			Field("project_identifier_id").
			Ref("projects").
			Unique().
			Required(),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("project_identifier_id").
			Unique(),
	}
}

// func (Project) Indexes() []ent.Index {
// 	return []ent.Index{
// 		index.Fields("ts").
// 			Annotations(entsql.IndexType("GIN")),
// 	}
// }
