package models

import "time"

type Project struct {
	ID               string
	Identifier       map[string][]string
	Name             map[string]string
	Description      map[string]string
	FoundingDate     string
	DissolutionDate  string
	Grant            string
	FundingProgramme string
	Acronym          string
	DateCreated      time.Time
	DateModified     time.Time
}
