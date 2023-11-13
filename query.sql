-- name: UpsertProject :exec
INSERT INTO projects(
        primary_identifier,
        identifiers,
        name,
        description,
        founding_date,
        dissolution_date,
        acronym,
        eu_acronym,
        eu_grant,
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
        $10,
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
    eu_acronym = excluded.eu_acronym,
    eu_grant = excluded.eu_grant,
    eu_funding_programme = excluded.eu_funding_programme,
    created_at = projects.created_at,
    updated_at = current_timestamp;

-- name: GetProject :one
SELECT pid,
    primary_identifier,
    identifiers,
    name,
    description,
    founding_date,
    dissolution_date,
    acronym,
    eu_acronym,
    eu_grant,
    eu_funding_programme,
    created_at,
    updated_at
FROM projects
WHERE primary_identifier = $1
LIMIT 1;

-- name: SuggestProjects :many

SELECT pid,
    primary_identifier,
    identifiers,
    name,
    description,
    founding_date,
    dissolution_date,
    acronym,
    eu_acronym,
    eu_grant,
    eu_funding_programme,
    created_at,
    updated_at
FROM projects
WHERE ts @@ websearch_to_tsquery($1)
LIMIT 10;

-- name: DeleteProject :one
DELETE FROM projects
WHERE primary_identifier = $1
RETURNING pid;