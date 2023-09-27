package models

import "time"

type Identifier struct {
	PropertyID string
	Value      string
}

type Grant struct {
	Identifier  string
	IsAwardedBy *FundingProgramme
}

type FundingProgramme struct {
	Name string
}

type Project struct {
	ID              string
	Identifier      []*Identifier
	HasAcronym      string
	IsFundedBy      *Grant
	Name            string
	Description     string
	FoundingDate    string
	DissolutionDate string
	DateCreated     *time.Time
	DateModified    *time.Time
}
