package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"entgo.io/ent/dialect"
	sqldialect "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ugent-library/projects/ent"
	"github.com/ugent-library/projects/ent/projectidentifier"
	"github.com/ugent-library/projects/models"
)

var ErrNotFound = errors.New("not found")
var ErrConstraint = errors.New("something went wrong")

type Repo struct {
	client *ent.Client
	config Config
}

type Config struct {
	Conn   string
	Secret []byte
}

func New(c Config) (*Repo, error) {
	db, err := sql.Open("pgx", c.Conn)
	if err != nil {
		return nil, err
	}

	driver := sqldialect.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver)).Debug()

	// err = client.Schema.Create(context.TODO(),
	// 	migrate.WithDropIndex(true),
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return &Repo{
		config: c,
		client: client,
	}, nil
}

func (r *Repo) AddProject(ctx context.Context, p *models.Project) error {
	sql := `
	WITH add_external_identifier AS (
		INSERT INTO project_identifiers (external_id)
		values ($1) ON CONFLICT (external_id) DO NOTHING
		RETURNING id
	),
	external_identifier AS (
		SELECT id
		FROM add_external_identifier
		UNION
		SELECT id
		FROM project_identifiers
		WHERE external_id = $1
	)
	INSERT INTO projects(
			project_identifier_id,
			identifier,
			name,
			description,
			founding_date,
			dissolution_date,
			acronym,
			grant_id,
			funding_programme,
			deleted,
			created_at,
			updated_at
		)
	SELECT external_identifier.id,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		current_timestamp,
		current_timestamp
	FROM external_identifier ON CONFLICT (project_identifier_id) DO
	UPDATE
	SET identifier = excluded.identifier,
		name = excluded.name,
		description = excluded.description,
		founding_date = excluded.founding_date,
		dissolution_date = excluded.dissolution_date,
		acronym = excluded.acronym,
		grant_id = excluded.grant_id,
		funding_programme = excluded.funding_programme,
		deleted = excluded.deleted,
		updated_at = excluded.updated_at
	WHERE projects.identifier != excluded.identifier
		OR projects.name != excluded.name
		OR projects.description != excluded.description
		OR projects.founding_date != excluded.founding_date
		OR projects.dissolution_date != excluded.dissolution_date
		OR projects.acronym != excluded.acronym
		OR projects.grant_id != excluded.grant_id
		OR projects.funding_programme != excluded.funding_programme
		OR projects.deleted != excluded.deleted
		OR projects.updated_at != excluded.updated_at
	`

	_, err := r.client.ExecContext(ctx, sql, p.ID, p.GetIdentifier(), p.GetName(), p.GetDescription(), p.FoundingDate, p.DissolutionDate, p.Acronym, p.Grant, p.FundingProgramme, p.Deleted)

	return err
}

func (r *Repo) GetProject(ctx context.Context, id string) (*models.Project, error) {
	row, err := r.client.ProjectIdentifier.Query().
		Where(projectidentifier.ExternalID(id)).
		QueryProjects().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, ErrNotFound
	}

	if ent.IsConstraintError(err) {
		return nil, ErrConstraint
	}

	if err != nil {
		return nil, err
	}

	return rowToProject(row), nil
}

func (r *Repo) DeleteProject(ctx context.Context, id string) error {
	p, err := r.client.ProjectIdentifier.Query().
		Where(projectidentifier.ExternalID(id)).
		QueryProjects().
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil
	case ent.IsConstraintError(err):
		return ErrConstraint
	case err != nil:
		return err
	}

	err = p.Update().
		SetDeleted(true).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) SuggestProjects(ctx context.Context, query string) ([]*models.Project, error) {
	tsQuery, tsQueryArgs := toTSQuery(query)
	tsQuery = "ts @@ " + tsQuery
	rows, err := r.client.Project.Query().Where(func(s *sqldialect.Selector) {
		s.Where(
			sqldialect.ExprP(tsQuery, tsQueryArgs...),
		)
	}).Limit(10).All(ctx)

	if err != nil {
		return nil, err
	}

	projects := make([]*models.Project, 0, len(rows))
	for _, row := range rows {
		projects = append(projects, rowToProject(row))
	}

	return projects, nil
}

func rowToProject(row *ent.Project) *models.Project {
	p := &models.Project{
		// ID:              row.ID,
		Identifier:      row.Identifier.Value,
		DateCreated:     row.CreatedAt,
		DateModified:    row.UpdatedAt,
		Name:            row.Name.Value,
		Description:     row.Description.Value,
		FoundingDate:    row.FoundingDate,
		DissolutionDate: row.DissolutionDate,
		Grant:           row.GrantID,
		Acronym:         row.Acronym,
	}

	return p
}

var regexMultipleSpaces = regexp.MustCompile(`\s+`)
var regexNoBrackets = regexp.MustCompile(`[\[\]()\{\}]`)

func toTSQuery(query string) (string, []any) {
	// remove duplicate spaces
	query = regexMultipleSpaces.ReplaceAllString(query, " ")
	// trim
	query = strings.TrimSpace(query)

	queryParts := make([]string, 0)
	queryArgs := make([]any, 0)
	argCounter := 0

	for _, qp := range strings.Split(query, " ") {
		// remove terms that contain brackets
		if regexNoBrackets.MatchString(qp) {
			continue
		}
		argCounter++

		// $1 || ':*'
		queryParts = append(queryParts, fmt.Sprintf("$%d || ':*'", argCounter))
		queryArgs = append(queryArgs, qp)
	}

	// $1:* & $2:*
	tsQuery := fmt.Sprintf(
		"to_tsquery('usimple', %s)",
		strings.Join(queryParts, " || ' & ' || "),
	)

	return tsQuery, queryArgs
}
