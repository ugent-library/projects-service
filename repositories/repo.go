package repositories

import (
	"context"
	"errors"
	"slices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ugent-library/projects-service/db"
	"github.com/ugent-library/projects-service/models"
)

var ErrNotFound = errors.New("not found")
var ErrConstraint = errors.New("something went wrong")

type Repo struct {
	config  Config
	db      *pgxpool.Pool
	queries *db.Queries
}

type Config struct {
	Conn string
}

func New(c Config) (*Repo, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, c.Conn)
	if err != nil {
		return nil, err
	}

	return &Repo{
		config:  c,
		db:      pool,
		queries: db.New(pool),
	}, nil
}

func (r *Repo) GetProject(ctx context.Context, id models.Identifier) (*models.ProjectRecord, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	project, err := r.queries.GetProject(ctx, db.GetProjectParams(id))
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	ids, err := r.queries.GetProjectIdentifiers(ctx, project.ID)
	if err != nil {
		return nil, err
	}

	p := &models.ProjectRecord{
		Project: models.Project{
			Name:            project.Name,
			Description:     project.Description,
			FoundingDate:    project.FoundingDate.String,
			DissolutionDate: project.DissolutionDate.String,
			Attributes:      project.Attributes,
			Identifiers:     make([]models.Identifier, len(ids)),
		},
		CreatedAt: project.CreatedAt.Time,
		UpdatedAt: project.UpdatedAt.Time,
	}

	for i, id := range ids {
		p.Project.Identifiers[i] = models.Identifier{Type: id.Type, Value: id.Value}
	}

	return p, tx.Commit(ctx)
}

func (r *Repo) AddProject(ctx context.Context, p *models.Project) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queries := r.queries.WithTx(tx)

	// Find existing projects

	var existingProjects []db.Project
	var existingIdentifiers []db.ProjectsIdentifier

	for _, id := range p.Identifiers {
		project, err := queries.GetProject(ctx, db.GetProjectParams(id))
		if err != nil && err != pgx.ErrNoRows {
			return err
		}
		if err == pgx.ErrNoRows {
			continue
		}

		if !slices.ContainsFunc(existingProjects, func(p db.Project) bool { return p.ID == project.ID }) {
			existingProjects = append(existingProjects, project)
			identifiers, err := queries.GetProjectIdentifiers(ctx, project.ID)
			if err != nil {
				return err
			}
			existingProjects = append(existingProjects, project)
			existingIdentifiers = append(existingIdentifiers, identifiers...)
		}
	}

	slices.SortFunc(existingProjects, func(a, b db.Project) int {
		if a.UpdatedAt.Time.Before(b.UpdatedAt.Time) {
			return 1
		}

		return -1
	})

	// Create a new project

	if len(existingProjects) == 0 {
		projectId, err := queries.CreateProject(ctx, db.CreateProjectParams{
			Name:            p.Name,
			Description:     p.Description,
			FoundingDate:    pgtype.Text{String: p.FoundingDate},
			DissolutionDate: pgtype.Text{String: p.DissolutionDate},
			Attributes:      p.Attributes,
		})
		if err != nil {
			return err
		}

		for _, id := range p.Identifiers {
			err := queries.CreateProjectIdentifier(ctx, db.CreateProjectIdentifierParams{
				ProjectID: projectId,
				Type:      id.Type,
				Value:     id.Value,
			})
			if err != nil {
				return err
			}
		}

		return tx.Commit(ctx)
	}

	// Update the most recent existing project

	projectID := existingProjects[0].ID

	err = queries.UpdateProject(ctx, db.UpdateProjectParams{
		ID:              projectID,
		Name:            p.Name,
		Description:     p.Description,
		FoundingDate:    pgtype.Text{String: p.FoundingDate},
		DissolutionDate: pgtype.Text{String: p.DissolutionDate},
		Attributes:      p.Attributes,
	})
	if err != nil {
		return err
	}

	for _, id := range existingIdentifiers {
		err = queries.DeleteProjectIdentifier(ctx, db.DeleteProjectIdentifierParams{
			Type:  id.Type,
			Value: id.Value,
		})
		if err != nil {
			return err
		}
	}

	for _, id := range p.Identifiers {
		err = queries.CreateProjectIdentifier(ctx, db.CreateProjectIdentifierParams{
			ProjectID: projectID,
			Type:      id.Type,
			Value:     id.Value,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// func (r *Repo) EachProject(ctx context.Context, fn func(p *models.Project) bool) error {
// 	var page int32
// 	var size int32

// 	page = 0
// 	size = 1000

// 	for {
// 		rows, err := r.queries.EachProject(ctx, db.EachProjectParams{Offset: page, Limit: size})
// 		if err != nil {
// 			return err
// 		}

// 		if len(rows) <= 0 {
// 			break
// 		}

// 		for _, row := range rows {
// 			pr := &models.Project{
// 				ID:              row.ExternalPrimaryIdentifier,
// 				Name:            row.Name,
// 				Description:     row.Description,
// 				Identifier:      row.ExternalIdentifiers,
// 				FoundingDate:    row.FoundingDate.String,
// 				DissolutionDate: row.DissolutionDate.String,
// 				Acronym:         row.Acronym,
// 				GrantCall:       row.EuGrantCall.String,
// 				DateCreated:     row.CreatedAt.Time,
// 				DateModified:    row.UpdatedAt.Time,
// 			}

// 			fn(pr)
// 		}

// 		page = page + size
// 	}

// 	return nil
// }

// func (r *Repo) EachProjectBetween(ctx context.Context, t1, t2 time.Time, fn func(p *models.Project) bool) error {
// 	rows, err := r.queries.BetweenProjects(ctx, db.BetweenProjectsParams{
// 		CreatedAt:   pgtype.Timestamptz{Time: t1, Valid: true},
// 		CreatedAt_2: pgtype.Timestamptz{Time: t2, Valid: true},
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	for _, row := range rows {
// 		pr := &models.Project{
// 			ID:              row.ExternalPrimaryIdentifier,
// 			Name:            row.Name,
// 			Description:     row.Description,
// 			Identifier:      row.ExternalIdentifiers,
// 			FoundingDate:    row.FoundingDate.String,
// 			DissolutionDate: row.DissolutionDate.String,
// 			Acronym:         row.Acronym,
// 			GrantCall:       row.EuGrantCall.String,
// 			DateCreated:     row.CreatedAt.Time,
// 			DateModified:    row.UpdatedAt.Time,
// 		}

// 		fn(pr)
// 	}

// 	return nil
// }
