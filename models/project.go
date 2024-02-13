package models

import (
	"time"
)

type Project struct {
	Name            []Translation `json:"name,omitempty"`
	Description     []Translation `json:"description,omitempty"`
	FoundingDate    string        `json:"founding_date,omitempty"`
	DissolutionDate string        `json:"dissolution_date,omitempty"`
	Deleted         bool          `json:"deleted,omitempty"`
	Attributes      []Attribute   `json:"attributes,omitempty"`
	Identifiers     []Identifier  `json:"identifiers,omitempty"`
}

type ProjectRecord struct {
	Project
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Identifier struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Attribute struct {
	Scope string `json:"scope,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Translation struct {
	Lang  string `json:"lang,omitempty"`
	Value string `json:"value,omitempty"`
}
