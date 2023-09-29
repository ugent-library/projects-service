package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

var timeUTC = func() time.Time {
	return time.Now().UTC()
}

type Project struct {
	ent.Schema
}

type Identifier struct {
	PropertyID string `json:"property_id"`
	Value      string `json:"value"`
}

type Grant struct {
	Identifier  string           `json:"identifier"`
	IsAwardedBy FundingProgramme `json:"isAwardedBy"`
}

type FundingProgramme struct {
	Name string `json:"name"`
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable().
			Unique().
			DefaultFunc(func() string {
				return ulid.Make().String()
			}),
		field.JSON("identifier", []Identifier{}).
			Default([]Identifier{}),
		field.String("name").
			Optional(),
		field.Text("description").
			Optional(),
		field.String("founding_date").
			Optional(),
		field.String("dissolution_date").
			Optional(),
		field.String("acronym").
			Optional(),
		field.String("grant").
			Optional(),
		field.String("funding_programme").
			Optional(),
		field.Time("created").
			Default(timeUTC).
			Immutable(),
		field.Time("modified").
			Default(timeUTC).
			UpdateDefault(timeUTC).
			Immutable(),
	}
}
