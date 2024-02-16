package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ugent-library/projects-service/models"
)

type Conn interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Begin(context.Context) (pgx.Tx, error)
}

type projectRow struct {
	ID              int64
	Name            []models.Translation
	Description     []models.Translation
	FoundingDate    pgtype.Text
	DissolutionDate pgtype.Text
	Attributes      []models.Attribute
	Identifiers     []models.Identifier
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
}

func (row projectRow) toProjectRecord() *models.ProjectRecord {
	return &models.ProjectRecord{
		Name:            row.Name,
		Description:     row.Description,
		FoundingDate:    row.FoundingDate.String,
		DissolutionDate: row.DissolutionDate.String,
		Attributes:      row.Attributes,
		Identifiers:     row.Identifiers,
		CreatedAt:       row.CreatedAt.Time,
		UpdatedAt:       row.UpdatedAt.Time,
	}
}

const getProjectQuery = `
WITH identifiers AS (
	SELECT pi1.*
	FROM projects_identifiers pi1
	LEFT JOIN projects_identifiers pi2 ON pi1.project_id = pi2.project_id
	WHERE pi2.type = $1 AND pi2.value = $2	
)
SELECT p.*, json_agg(json_build_object('type', i.type, 'value', i.value)) AS identifiers
FROM projects p, identifiers i WHERE p.id = i.project_id
GROUP BY p.id;
`

const getEachProjectQuery = `
SELECT p.*, json_agg(json_build_object('type', pi.type, 'value', pi.value)) AS identifiers
FROM projects p
LEFT JOIN projects_identifiers pi ON p.id = pi.project_id
GROUP BY p.id;
`
