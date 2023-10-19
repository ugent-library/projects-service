package api

import (
	"context"
	"fmt"

	"github.com/go-faster/errors"
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

func (s *Service) AddProject(ctx context.Context, req *AddProject) (AddProjectRes, error) {
	p := &models.Project{}

	if v, ok := req.GetID().Get(); ok {
		tmp, err := s.repo.GetProject(ctx, v)
		if !errors.Is(err, repositories.ErrNotFound) {
			p = tmp
		}
	}

	if v, ok := req.GetID().Get(); ok {
		p.ID = v
	}

	if v, ok := req.GetCreated().Get(); ok {
		p.DateCreated = v
	}

	if v, ok := req.GetModified().Get(); ok {
		p.DateModified = v
	}

	if ids := req.GetIdentifier(); len(ids) > 0 {
		tmp := make(map[string][]string)
		for _, id := range ids {
			tmp[id.GetPropertyID()] = append(tmp[id.GetPropertyID()], id.GetValue())
		}
		p.Identifier = tmp
	}

	if strs := req.GetName(); len(strs) > 0 {
		tmp := make(map[string]string)
		for _, str := range strs {
			tmp[str.GetLanguage()] = str.GetValue()
		}

		p.Name = tmp
	}

	if v, ok := req.GetDescription().Get(); ok {
		p.Description = v
	}

	if v, ok := req.GetFoundingDate().Get(); ok {
		p.FoundingDate = v
	}

	if v, ok := req.GetDissolutionDate().Get(); ok {
		p.DissolutionDate = v
	}

	if v, ok := req.GetHasAcronym().Get(); ok {
		p.Acronym = v
	}

	if fb, ok := req.GetIsFundedBy().Get(); ok {
		id := fb.GetIdentifier()
		p.Grant = id

		if ab, ok := fb.GetIsAwardedBy().Get(); ok {
			name := ab.GetName()
			p.FundingProgramme = name
		}
	}

	if err := s.repo.AddProject(ctx, p); err != nil {
		return nil, err
	}

	return &AddProjectOK{}, nil
}

func (s *Service) DeleteProject(ctx context.Context, req *DeleteProjectRequest) (DeleteProjectRes, error) {
	v := req.GetID()
	err := s.repo.DeleteProject(ctx, v)
	if errors.Is(err, repositories.ErrNotFound) {
		return &ErrorStatusCode{
			StatusCode: 404,
			Response: Error{
				Code:    404,
				Message: fmt.Sprintf("Project not found: %s", v),
			},
		}, nil
	}

	return &DeleteProjectOK{}, nil
}

func (s *Service) GetProject(ctx context.Context, req *GetProjectRequest) (GetProjectRes, error) {
	v := req.GetID()
	p, err := s.repo.GetProject(ctx, v)
	if errors.Is(err, repositories.ErrNotFound) {
		return &ErrorStatusCode{
			StatusCode: 404,
			Response: Error{
				Code:    404,
				Message: fmt.Sprintf("Project not found: %s", v),
			},
		}, nil
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
		Data: make([]GetProject, 0, len(ps)),
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

func mapToOASProject(p *models.Project) *GetProject {
	sids := make([]GetProjectIdentifierItem, 0)
	for prop, ids := range p.Identifier {
		for _, id := range ids {
			sids = append(sids, GetProjectIdentifierItem{
				Type:       "PropertyValue",
				PropertyID: prop,
				Value:      id,
			})
		}
	}

	r := &GetProject{
		Type:       "ResearchProject",
		Identifier: sids,
		Created:    p.DateCreated,
		Modified:   p.DateModified,
	}

	r.ID = p.ID

	strs := make([]GetProjectNameItem, 0)
	for lang, val := range p.Name {
		strs = append(strs, GetProjectNameItem{
			Language: lang,
			Value:    val,
		})
	}

	if p.Description != "" {
		r.SetDescription(NewOptString(p.Description))
	}

	if p.FoundingDate != "" {
		r.SetFoundingDate(NewOptString(p.FoundingDate))
	}
	if p.DissolutionDate != "" {
		r.SetDissolutionDate(NewOptString(p.DissolutionDate))
	}

	if p.Acronym != "" {
		r.SetHasAcronym(NewOptString(p.Acronym))
	}

	if p.Grant != "" {
		g := GetProjectIsFundedBy{
			Type:       "Grant",
			Identifier: p.Grant,
		}

		if p.FundingProgramme != "" {
			g.IsAwardedBy.SetTo(GetProjectIsFundedByIsAwardedBy{
				Type: "FundingProgramme",
				Name: p.FundingProgramme,
			})
		}

		r.IsFundedBy.SetTo(g)
	}

	return r
}
