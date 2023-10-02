package models

import "time"

type Identifier struct {
	PropertyID string
	Value      string
}

type Project struct {
	ID               string
	Identifier       []*Identifier
	Name             *string
	Description      *string
	FoundingDate     *string
	DissolutionDate  *string
	Grant            *string
	FundingProgramme *string
	Acronym          *string
	DateCreated      time.Time
	DateModified     time.Time
}
