package api

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/ugent-library/projects-service/models"
	"github.com/ugent-library/projects-service/repositories"
)

type Service struct {
	repo *repositories.Repo
	// searcher search.Searcher
}

func NewService(repo *repositories.Repo) *Service {
	return &Service{
		repo: repo,
		//	searcher: searcher,
	}
}

func (s *Service) GetProject(ctx context.Context, id *Identifier) (GetProjectRes, error) {
	p, err := s.repo.GetProject(ctx, models.Identifier(*id))
	if errors.Is(err, repositories.ErrNotFound) {
		return nil, &ErrorStatusCode{
			StatusCode: 404,
			Response: Error{
				Code:    404,
				Message: "Project not found",
			},
		}
	}
	if err != nil {
		return nil, err
	}

	attributes := make([]Attribute, len(p.Attributes))
	for i, attr := range p.Attributes {
		attributes[i] = Attribute(attr)
	}

	identifiers := make([]Identifier, len(p.Identifiers))
	for i, id := range p.Identifiers {
		identifiers[i] = Identifier(id)
	}

	names := make([]Translation, len(p.Name))
	for i, name := range p.Name {
		names[i] = Translation(name)
	}

	descriptions := make([]Translation, len(p.Description))
	for i, desc := range p.Description {
		descriptions[i] = Translation(desc)
	}

	return &ProjectRecord{
		Name:            names,
		Description:     descriptions,
		FoundingDate:    NewOptString(p.FoundingDate),
		DissolutionDate: NewOptString(p.DissolutionDate),
		Attributes:      attributes,
		Identifiers:     identifiers,
	}, nil
}

func (s *Service) AddProject(ctx context.Context, p *Project) error {
	attributes := make([]models.Attribute, len(p.Attributes))
	for i, attr := range p.Attributes {
		attributes[i] = models.Attribute(attr)
	}

	identifiers := make([]models.Identifier, len(p.Identifiers))
	for i, id := range p.Identifiers {
		identifiers[i] = models.Identifier(id)
	}

	names := make([]models.Translation, len(p.Name))
	for i, name := range p.Name {
		names[i] = models.Translation(name)
	}

	descriptions := make([]models.Translation, len(p.Description))
	for i, desc := range p.Description {
		descriptions[i] = models.Translation(desc)
	}

	foundingDate := ""
	if v, ok := p.GetFoundingDate().Get(); ok {
		foundingDate = v
	}

	dissolutionDate := ""
	if v, ok := p.GetDissolutionDate().Get(); ok {
		dissolutionDate = v
	}

	return s.repo.AddProject(ctx, &models.Project{
		Name:            names,
		Description:     descriptions,
		FoundingDate:    foundingDate,
		DissolutionDate: dissolutionDate,
		Attributes:      attributes,
		Identifiers:     identifiers,
	})
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
