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
		Name:            req.GetName().Or(""),
		Description:     req.GetDescription().Or(""),
		Identifier:      ids,
		FoundingDate:    req.GetFoundingDate().Or(""),
		DissolutionDate: req.GetDissolutionDate().Or(""),
		Acronym:         req.GetHasAcronym().Or(""),
	}

	if fb, ok := req.GetIsFundedBy().Get(); ok {
		p.Grant = fb.GetIdentifier()

		if ab, ok := fb.GetIsAwardedBy().Get(); ok {
			p.FundingProgramme = ab.Name
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

	g := NewOptGetProjectResponseIsFundedBy(GetProjectResponseIsFundedBy{})
	if p.Grant != "" {
		g.Value.SetIdentifier(p.Grant)
		g.Value.SetType("Grant")

		fp := NewOptGetProjectResponseIsFundedByIsAwardedBy(GetProjectResponseIsFundedByIsAwardedBy{})

		if p.FundingProgramme != "" {
			fp.Value.SetName(p.FundingProgramme)
			fp.Value.SetType("FundingProgramme")
		}
	} else {
		g.Reset()
	}

	return &GetProjectResponse{
		Type:            "ResearchProject",
		Identifier:      ids,
		Name:            NewOptString(p.Name),
		Description:     NewOptString(p.Description),
		IsFundedBy:      g,
		HasAcronym:      NewOptString(p.Acronym),
		FoundingDate:    NewOptString(p.FoundingDate),
		DissolutionDate: NewOptString(p.DissolutionDate),
		Created:         *p.DateCreated,
		Modified:        *p.DateModified,
	}, nil
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
