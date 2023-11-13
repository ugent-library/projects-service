package repositories

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ugent-library/projects/models"
	"github.com/ugent-library/projects/sqlc"
)

var ErrNotFound = errors.New("not found")
var ErrConstraint = errors.New("something went wrong")

type Repo struct {
	client *sqlc.Queries
	config Config
}

type Config struct {
	Conn   string
	Secret []byte
}

func New(c Config) (*Repo, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, c.Conn)
	if err != nil {
		return nil, err
	}

	client := sqlc.New(conn)

	return &Repo{
		config: c,
		client: client,
	}, nil
}

func (r *Repo) AddProject(ctx context.Context, p *models.Project) error {
	d := sqlc.UpsertProjectParams{
		PrimaryIdentifier: "test",
		Identifiers:       p.Identifier,
		Name:              p.Name,
		Description:       p.Description,
		FoundingDate:      pgtype.Text{String: p.FoundingDate, Valid: true},
		DissolutionDate:   pgtype.Text{String: p.DissolutionDate, Valid: true},
		Acronym:           pgtype.Text{String: p.Acronym, Valid: true},
		// TODO eu_acronym
		// TODO eu_grant
		// TODO eu_funding_programme
	}

	err := r.client.UpsertProject(ctx, d)

	return err
}

func (r *Repo) GetProject(ctx context.Context, id string) (*models.Project, error) {
	row, err := r.client.GetProject(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	p := &models.Project{
		ID:              row.PrimaryIdentifier,
		Name:            row.Name,
		Description:     row.Description,
		Identifier:      row.Identifiers,
		FoundingDate:    row.FoundingDate.String,
		DissolutionDate: row.DissolutionDate.String,
		Acronym:         row.Acronym.String,
		// Grant
		// FundingProgramme
		DateCreated:  row.CreatedAt.Time,
		DateModified: row.UpdatedAt.Time,
	}

	return p, nil
}

func (r *Repo) DeleteProject(ctx context.Context, id string) error {
	_, err := r.client.DeleteProject(ctx, id)

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return ErrNotFound
	case err != nil:
		return err
	}

	return nil
}

func (r *Repo) SuggestProjects(ctx context.Context, query string) ([]*models.Project, error) {
	toTSQuery := toTSQuery(query)
	rows, err := r.client.SuggestProjects(ctx, toTSQuery)

	if err != nil {
		return nil, err
	}

	projects := make([]*models.Project, 0, len(rows))
	if rows == nil {
		return projects, nil
	}

	for _, row := range rows {
		projects = append(projects, &models.Project{
			ID:              row.PrimaryIdentifier,
			Name:            row.Name,
			Description:     row.Description,
			Identifier:      row.Identifiers,
			FoundingDate:    row.FoundingDate.String,
			DissolutionDate: row.DissolutionDate.String,
			Acronym:         row.Acronym.String,
			// Grant
			// FundingProgramme
			DateCreated:  row.CreatedAt.Time,
			DateModified: row.UpdatedAt.Time,
		})
	}
	return projects, nil
}

var regexNoMultipleSpaces = regexp.MustCompile(`\s+`)
var regexNoBrackets = regexp.MustCompile(`[\[\]()\{\}]`)

func toTSQuery(query string) string {
	query = regexNoMultipleSpaces.ReplaceAllString(query, " ")
	query = strings.TrimSpace(query)

	parts := make([]string, 0)
	for _, qp := range strings.Split(query, " ") {
		if regexNoBrackets.MatchString(qp) {
			continue
		}

		parts = append(parts, fmt.Sprintf("'%s' || ':*'", qp))
	}

	d := fmt.Sprintf(
		"%s",
		strings.Join(parts, " || ' & ' || "),
	)

	return d
}

// func toTSQuery2(query string) (string, []any) {
// 	// remove duplicate spaces
// 	query = regexMultipleSpaces.ReplaceAllString(query, " ")
// 	// trim
// 	query = strings.TrimSpace(query)

// 	queryParts := make([]string, 0)
// 	queryArgs := make([]any, 0)
// 	argCounter := 0

// 	for _, qp := range strings.Split(query, " ") {
// 		// remove terms that contain brackets
// 		if regexNoBrackets.MatchString(qp) {
// 			continue
// 		}
// 		argCounter++

// 		// $1 || ':*'
// 		queryParts = append(queryParts, fmt.Sprintf("$%d || ':*'", argCounter))
// 		queryArgs = append(queryArgs, qp)
// 	}

// 	// $1:* & $2:*
// 	tsQuery := fmt.Sprintf(
// 		"to_tsquery('usimple', %s)",
// 		strings.Join(queryParts, " || ' & ' || "),
// 	)

// 	return tsQuery, queryArgs
// }
