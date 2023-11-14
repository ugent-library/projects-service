package models

import (
	"time"
)

type Identifiers struct {
	Value map[string][]string `json:"value"`
}

type TranslatedString struct {
	Value map[string]string `json:"value"`
}

type Acronym struct {
	Value []string `json:value`
}

type Project struct {
	ID               string
	Identifier       Identifiers
	Name             TranslatedString
	Description      TranslatedString
	FoundingDate     string
	DissolutionDate  string
	GrantCall        string
	FundingProgramme string
	Acronym          Acronym
	Deleted          bool
	DateCreated      time.Time
	DateModified     time.Time
}
