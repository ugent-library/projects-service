package repositories

import (
	"context"
	"database/sql"
	"errors"

	"entgo.io/ent/dialect"
	sqldialect "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ugent-library/projects/ent"
	"github.com/ugent-library/projects/ent/migrate"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
	"github.com/ugent-library/projects/models"
)

var ErrNotFound = errors.New("not found")

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

	err = client.Schema.Create(context.TODO(),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}

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
		SetName(p.Name).
		SetDescription(p.Description).
		SetIdentifier(sids).
		SetFoundingDate(p.FoundingDate).
		SetDissolutionDate(p.DissolutionDate).
		Save(ctx)

	return err
}

func (r *Repo) GetProject(ctx context.Context, id string) (*models.Project, error) {
	row, err := r.client.Project.Query().
		Where(project.IDEQ(id)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	ids := make([]*models.Identifier, 0, len(row.Identifier))
	for _, id := range row.Identifier {
		ids = append(ids, &models.Identifier{
			PropertyID: id.PropertyID,
			Value:      id.Value,
		})
	}

	p := &models.Project{
		ID:              row.ID,
		Name:            row.Name,
		Description:     row.Description,
		Identifier:      ids,
		FoundingDate:    row.FoundingDate,
		DissolutionDate: row.DissolutionDate,
		DateCreated:     &row.Created,
		DateModified:    &row.Modified,
	}

	return p, nil
}
