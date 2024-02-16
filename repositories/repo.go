package repositories

import (
	"context"
	"errors"
	"slices"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ugent-library/projects-service/db"
	"github.com/ugent-library/projects-service/models"
)

var ErrNotFound = errors.New("not found")
var ErrConstraint = errors.New("something went wrong")

type Repo struct {
	conn    Conn
	queries *db.Queries
}

type RepoConfig struct {
	Conn Conn
}

func New(c RepoConfig) (*Repo, error) {
	return &Repo{
		conn:    c.Conn,
		queries: db.New(c.Conn),
	}, nil
}

func (r *Repo) GetProject(ctx context.Context, id models.Identifier) (*models.ProjectRecord, error) {
	var row projectRow
	err := pgxscan.Get(ctx, r.conn, &row, getProjectQuery, id.Type, id.Value)
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return row.toProjectRecord(), nil
}

func (r *Repo) AddProject(ctx context.Context, p *models.Project) error {
	tx, err := r.conn.Begin(ctx)
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

func (r *Repo) EachProject(ctx context.Context, fn func(*models.ProjectRecord) bool) error {
	rows, err := r.conn.Query(ctx, getEachProjectQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	rs := pgxscan.NewRowScanner(rows)
	for rows.Next() {
		var row projectRow
		if err := rs.Scan(&row); err != nil {
			return err
		}

		if ok := fn(row.toProjectRecord()); !ok {
			break
		}
	}

	return rows.Err()
}
