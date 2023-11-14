// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	models "github.com/ugent-library/projects/models"
)

const deleteProject = `-- name: DeleteProject :one
DELETE FROM projects
WHERE primary_identifier = $1
RETURNING pid
`

func (q *Queries) DeleteProject(ctx context.Context, primaryIdentifier string) (int64, error) {
	row := q.db.QueryRow(ctx, deleteProject, primaryIdentifier)
	var pid int64
	err := row.Scan(&pid)
	return pid, err
}

const getProject = `-- name: GetProject :one
SELECT pid,
    primary_identifier,
    identifiers,
    name,
    description,
    founding_date,
    dissolution_date,
    acronym,
    eu_grant_call,
    eu_funding_programme,
    created_at,
    updated_at
FROM projects
WHERE primary_identifier = $1
LIMIT 1
`

type GetProjectRow struct {
	Pid                int64
	PrimaryIdentifier  string
	Identifiers        models.Identifiers
	Name               models.TranslatedString
	Description        models.TranslatedString
	FoundingDate       pgtype.Text
	DissolutionDate    pgtype.Text
	Acronym            models.Acronym
	EuGrantCall        pgtype.Text
	EuFundingProgramme pgtype.Text
	CreatedAt          pgtype.Timestamptz
	UpdatedAt          pgtype.Timestamptz
}

func (q *Queries) GetProject(ctx context.Context, primaryIdentifier string) (GetProjectRow, error) {
	row := q.db.QueryRow(ctx, getProject, primaryIdentifier)
	var i GetProjectRow
	err := row.Scan(
		&i.Pid,
		&i.PrimaryIdentifier,
		&i.Identifiers,
		&i.Name,
		&i.Description,
		&i.FoundingDate,
		&i.DissolutionDate,
		&i.Acronym,
		&i.EuGrantCall,
		&i.EuFundingProgramme,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const suggestProjects = `-- name: SuggestProjects :many

SELECT pid,
    primary_identifier,
    identifiers,
    name,
    description,
    founding_date,
    dissolution_date,
    acronym,
    eu_grant_call,
    eu_funding_programme,
    created_at,
    updated_at
FROM projects
WHERE ts @@ websearch_to_tsquery($1)
LIMIT 10
`

type SuggestProjectsRow struct {
	Pid                int64
	PrimaryIdentifier  string
	Identifiers        models.Identifiers
	Name               models.TranslatedString
	Description        models.TranslatedString
	FoundingDate       pgtype.Text
	DissolutionDate    pgtype.Text
	Acronym            models.Acronym
	EuGrantCall        pgtype.Text
	EuFundingProgramme pgtype.Text
	CreatedAt          pgtype.Timestamptz
	UpdatedAt          pgtype.Timestamptz
}

func (q *Queries) SuggestProjects(ctx context.Context, websearchToTsquery string) ([]SuggestProjectsRow, error) {
	rows, err := q.db.Query(ctx, suggestProjects, websearchToTsquery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SuggestProjectsRow
	for rows.Next() {
		var i SuggestProjectsRow
		if err := rows.Scan(
			&i.Pid,
			&i.PrimaryIdentifier,
			&i.Identifiers,
			&i.Name,
			&i.Description,
			&i.FoundingDate,
			&i.DissolutionDate,
			&i.Acronym,
			&i.EuGrantCall,
			&i.EuFundingProgramme,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertProject = `-- name: UpsertProject :exec
INSERT INTO projects(
        primary_identifier,
        identifiers,
        name,
        description,
        founding_date,
        dissolution_date,
        acronym,
        eu_grant_call,
        eu_funding_programme,
        created_at,
        updated_at
    )
VALUES(
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        current_timestamp,
        current_timestamp
    ) ON CONFLICT (primary_identifier) DO
UPDATE
SET identifiers = excluded.identifiers,
    name = excluded.name,
    description = excluded.description,
    founding_date = excluded.founding_date,
    dissolution_date = excluded.dissolution_date,
    acronym = excluded.acronym,
    eu_grant_call = excluded.eu_grant_call,
    eu_funding_programme = excluded.eu_funding_programme,
    created_at = projects.created_at,
    updated_at = current_timestamp
`

type UpsertProjectParams struct {
	PrimaryIdentifier  string
	Identifiers        models.Identifiers
	Name               models.TranslatedString
	Description        models.TranslatedString
	FoundingDate       pgtype.Text
	DissolutionDate    pgtype.Text
	Acronym            models.Acronym
	EuGrantCall        pgtype.Text
	EuFundingProgramme pgtype.Text
}

func (q *Queries) UpsertProject(ctx context.Context, arg UpsertProjectParams) error {
	_, err := q.db.Exec(ctx, upsertProject,
		arg.PrimaryIdentifier,
		arg.Identifiers,
		arg.Name,
		arg.Description,
		arg.FoundingDate,
		arg.DissolutionDate,
		arg.Acronym,
		arg.EuGrantCall,
		arg.EuFundingProgramme,
	)
	return err
}
