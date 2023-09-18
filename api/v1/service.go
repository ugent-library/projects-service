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
	for _, id := range req.Identifier {
		ids = append(ids, &models.Identifier{
			PropertyID: id.PropertyID,
			Value:      id.Value,
		})
	}

	return s.repo.AddProject(ctx, &models.Project{
		Name:            req.Name,
		Description:     req.Description,
		Identifier:      ids,
		FoundingDate:    req.FoundingDate,
		DissolutionDate: req.DissolutionDate,
	})
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

	return &GetProjectResponse{
		Type:            "ResearchProject",
		Identifier:      ids,
		Name:            p.Name,
		FoundingDate:    p.FoundingDate,
		DissolutionDate: p.DissolutionDate,
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
