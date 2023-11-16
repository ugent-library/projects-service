package models

import (
	"time"
)

type ExternalIdentifiers = map[string][]string
type TranslatedString = map[string]string
type Acronym = []string

type Project struct {
	ID               string
	Identifier       ExternalIdentifiers
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
