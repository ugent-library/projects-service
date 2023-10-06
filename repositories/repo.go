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
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
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
	client := ent.NewClient(ent.Driver(driver))

	return &Repo{
		config: c,
		client: client,
	}, nil
}

func (r *Repo) AddProject(ctx context.Context, p *models.Project) error {
	sids := make([]schema.Identifier, 0, len(p.Identifier))
	for _, id := range p.Identifier {
		sids = append(sids, schema.Identifier{
			PropertyID: id.PropertyID,
			Value:      id.Value,
		})
	}

	_, err := r.client.Project.Create().
		SetIdentifier(sids).
		SetNillableName(p.Name).
		SetNillableDescription(p.Description).
		SetNillableFoundingDate(p.FoundingDate).
		SetNillableDissolutionDate(p.DissolutionDate).
		SetNillableGrant(p.Grant).
		SetNillableFundingProgramme(p.FundingProgramme).
		SetNillableAcronym(p.Acronym).
		Save(ctx)

	return err
}

func (r *Repo) GetProject(ctx context.Context, id string) (*models.Project, error) {
	row, err := r.client.Project.Query().
		Where(project.IDEQ(id)).
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
	ids := make([]*models.Identifier, 0, len(row.Identifier))
	for _, id := range row.Identifier {
		ids = append(ids, &models.Identifier{
			PropertyID: id.PropertyID,
			Value:      id.Value,
		})
	}

	p := &models.Project{
		ID:               row.ID,
		Identifier:       ids,
		Name:             row.Name,
		Description:      row.Description,
		FoundingDate:     row.FoundingDate,
		DissolutionDate:  row.DissolutionDate,
		FundingProgramme: row.FundingProgramme,
		Grant:            row.Grant,
		Acronym:          row.Acronym,
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
