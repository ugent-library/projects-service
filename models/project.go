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

type Project struct {
	ID               string
	Identifier       Identifiers
	Name             TranslatedString
	Description      TranslatedString
	FoundingDate     string
	DissolutionDate  string
	Grant            string
	FundingProgramme string
	Acronym          string
	Deleted          bool
	DateCreated      time.Time
	DateModified     time.Time
}
