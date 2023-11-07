package models

import (
	"encoding/json"
	"time"
)

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
	Deleted          bool
	DateCreated      time.Time
	DateModified     time.Time
}

func (p *Project) GetIdentifier() string {
	d, err := json.Marshal(struct {
		Value map[string][]string `json:"value"`
	}{
		Value: p.Identifier,
	})

	if err != nil {
		panic(err)
	}

	return string(d)
}

func (p *Project) GetName() string {
	d, err := json.Marshal(struct {
		Value map[string]string `json:"value"`
	}{
		Value: p.Name,
	})

	if err != nil {
		panic(err)
	}

	return string(d)
}

func (p *Project) GetDescription() string {
	d, err := json.Marshal(struct {
		Value map[string]string `json:"value"`
	}{
		Value: p.Description,
	})

	if err != nil {
		panic(err)
	}

	return string(d)
}
