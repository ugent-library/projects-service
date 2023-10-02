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

func (s *Service) AddProject(ctx context.Context, req *AddProjectRequest) error {
	ids := make([]*models.Identifier, 0, len(req.Identifier))
	for _, id := range req.GetIdentifier() {
		ids = append(ids, &models.Identifier{
			PropertyID: id.GetPropertyID(),
			Value:      id.GetValue(),
		})
	}

	p := &models.Project{
		Identifier: ids,
	}

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

func (s *Service) GetProject(ctx context.Context, req *GetProjectRequest) (*GetProjectResponse, error) {
	p, err := s.repo.GetProject(ctx, req.ID)

	if err != nil {
		return nil, err
	}

	ids := make([]GetProjectResponseIdentifierItem, 0, len(p.Identifier))
	for _, id := range p.Identifier {
		ids = append(ids, GetProjectResponseIdentifierItem{
			Type:       "PropertyValue",
			PropertyID: id.PropertyID,
			Value:      id.Value,
		})
	}

	r := &GetProjectResponse{
		Type:       "ResearchProject",
		Identifier: ids,
		Created:    p.DateCreated,
		Modified:   p.DateModified,
	}

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

	r.IsFundedBy.SetTo(GetProjectResponseIsFundedBy{})
	r.IsFundedBy.SetToNull()
	if p.Grant != nil {
		g := GetProjectResponseIsFundedBy{
			Type:       "Grant",
			Identifier: *p.Grant,
		}

		g.IsAwardedBy.SetTo(GetProjectResponseIsFundedByIsAwardedBy{})
		g.IsAwardedBy.SetToNull()
		if p.FundingProgramme != nil {
			g.IsAwardedBy.SetTo(GetProjectResponseIsFundedByIsAwardedBy{
				Type: "FundingProgramme",
				Name: *p.FundingProgramme,
			})
		}

		r.IsFundedBy.SetTo(g)
	}

	return r, nil
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
