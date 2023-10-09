package api

import (
	"context"

	"github.com/ugent-library/projects/models"
	"github.com/ugent-library/projects/repositories"
)

type Service struct {
	repo *repositories.Repo
}

func NewService(repo *repositories.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddProject(ctx context.Context, req *Project) error {
	p := &models.Project{}

	if v, ok := req.ID.Get(); ok {
		p.ID = v
	}

	if v, ok := req.Created.Get(); ok {
		p.DateCreated = v
	}

	if v, ok := req.Modified.Get(); ok {
		p.DateModified = v
	}

	ids := make(map[string][]string)
	for _, id := range req.GetIdentifier() {
		ids[id.GetPropertyID()] = append(ids[id.GetPropertyID()], id.GetValue())
	}

	p.Identifier = ids

	if v, ok := req.GetName().Get(); ok {
		p.Name = &v
	}

	if v, ok := req.GetDescription().Get(); ok {
		p.Description = &v
	}

	if v, ok := req.GetFoundingDate().Get(); ok {
		p.FoundingDate = &v
	}

	if v, ok := req.GetDissolutionDate().Get(); ok {
		p.DissolutionDate = &v
	}

	if v, ok := req.GetHasAcronym().Get(); ok {
		p.Acronym = &v
	}

	if fb, ok := req.GetIsFundedBy().Get(); ok {
		id := fb.GetIdentifier()
		p.Grant = &id

		if ab, ok := fb.GetIsAwardedBy().Get(); ok {
			name := ab.GetName()
			p.FundingProgramme = &name
		}
	}

	return s.repo.AddProject(ctx, p)
}

func (s *Service) GetProject(ctx context.Context, req *GetProjectRequest) (*Project, error) {
	p, err := s.repo.GetProject(ctx, req.ID)

	if err != nil {
		return nil, err
	}

	oasp := mapToOASProject(p)

	return oasp, nil
}

func (s *Service) SuggestProjects(ctx context.Context, req *SuggestProjectsRequest) (*SuggestProjectsResponse, error) {
	ps, err := s.repo.SuggestProjects(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	res := &SuggestProjectsResponse{
		Data: make([]Project, 0, len(ps)),
	}

	for _, p := range ps {
		res.Data = append(res.Data, *mapToOASProject(p))
	}

	return res, nil
}

func (s *Service) NewError(ctx context.Context, err error) *ErrorStatusCode {
	return &ErrorStatusCode{
		StatusCode: 500,
		Response: Error{
			Code:    500,
			Message: err.Error(),
		},
	}
}

func mapToOASProject(p *models.Project) *Project {
	sids := make([]ProjectIdentifierItem, 0)
	for k, ids := range p.Identifier {
		for _, id := range ids {
			sids = append(sids, ProjectIdentifierItem{
				Type:       "PropertyValue",
				PropertyID: k,
				Value:      id,
			})
		}
	}

	r := &Project{
		Type:       "ResearchProject",
		Identifier: sids,
		Created:    NewOptDateTime(p.DateCreated),
		Modified:   NewOptDateTime(p.DateModified),
	}

	r.ID.SetTo(p.ID)

	r.Name = NewNilString("")
	r.Name.SetToNull()
	if p.Name != nil {
		r.Name.SetTo(*p.Name)
	}

	r.Description = NewNilString("")
	r.Description.SetToNull()
	if p.Description != nil {
		r.Description.SetTo(*p.Description)
	}

	r.FoundingDate = NewNilString("")
	r.FoundingDate.SetToNull()
	if p.FoundingDate != nil {
		r.FoundingDate.SetTo(*p.FoundingDate)
	}

	r.DissolutionDate = NewNilString("")
	r.DissolutionDate.SetToNull()
	if p.DissolutionDate != nil {
		r.DissolutionDate.SetTo(*p.DissolutionDate)
	}

	r.HasAcronym = NewNilString("")
	r.HasAcronym.SetToNull()
	if p.Acronym != nil {
		r.HasAcronym.SetTo(*p.Acronym)
	}

	r.IsFundedBy.SetTo(ProjectIsFundedBy{})
	r.IsFundedBy.SetToNull()
	if p.Grant != nil {
		g := ProjectIsFundedBy{
			Type:       "Grant",
			Identifier: *p.Grant,
		}

		g.IsAwardedBy.SetTo(ProjectIsFundedByIsAwardedBy{})
		g.IsAwardedBy.SetToNull()
		if p.FundingProgramme != nil {
			g.IsAwardedBy.SetTo(ProjectIsFundedByIsAwardedBy{
				Type: "FundingProgramme",
				Name: *p.FundingProgramme,
			})
		}

		r.IsFundedBy.SetTo(g)
	}

	return r
}
