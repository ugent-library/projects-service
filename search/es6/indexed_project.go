package es6

import (
	"log"

	"github.com/ugent-library/projects-service/models"
)

type projectDocument struct {
	DateCreated          string   `json:"date_created,omitempty"`
	DateUpdated          string   `json:"date_updated,omitempty"`
	Name                 []string `json:"name,omitempty"`
	Acronym              []string `json:"acronym,omitempty"`
	EUCallID             string   `json:"eu_call_id,omitempty"`
	EUFrameworkProgramme string   `json:"eu_framework_programme,omitempty"`
	CordisID             []string `json:"cordis_id,omitempty"`
	IwetoID              []string `json:"iweto_id,omitempty"`
	GismoID              []string `json:"gismo_id,omitempty"`
}

func NewProjectDocument(p *models.Project) *projectDocument {
	//time.RFC3339 does not include milliseconds
	const TimeFormatUTC = "2006-01-02T15:04:05.999Z"

	doc := &projectDocument{
		DateCreated:          p.DateCreated.UTC().Format(TimeFormatUTC),
		DateUpdated:          p.DateModified.UTC().Format(TimeFormatUTC),
		EUCallID:             p.GrantCall,
		EUFrameworkProgramme: p.FundingProgramme,
	}

	log.Println("----")
	log.Println(p.DateCreated)
	log.Println("----")

	tmp := make([]string, len(p.Name))
	for _, v := range p.Name {
		tmp = append(tmp, v)
	}
	doc.Name = tmp

	acrs := make([]string, len(p.Acronym))
	acrs = append(acrs, p.Acronym...)
	doc.Acronym = acrs

	for k, v := range p.Identifier {
		switch k {
		case "CORDIS":
			doc.CordisID = v
		case "IWETO":
			doc.IwetoID = v
		case "GISMO":
			doc.GismoID = v
		}
	}
	return doc
}
