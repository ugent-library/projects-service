-- name: UpsertProject :exec
INSERT INTO projects(
        external_primary_identifier,
        external_identifiers,
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
    ) ON CONFLICT (external_primary_identifier) DO
UPDATE
SET external_identifiers = excluded.external_identifiers,
    name = excluded.name,
    description = excluded.description,
    founding_date = excluded.founding_date,
    dissolution_date = excluded.dissolution_date,
    acronym = excluded.acronym,
    eu_grant_call = excluded.eu_grant_call,
    eu_funding_programme = excluded.eu_funding_programme,
    created_at = projects.created_at,
    updated_at = current_timestamp;

-- name: GetProject :one
SELECT pk,
    external_primary_identifier,
    external_identifiers,
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
WHERE external_primary_identifier = $1
LIMIT 1;

-- name: SuggestProjects :many

SELECT pk,
    external_primary_identifier,
    external_identifiers,
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
WHERE ts @@ to_tsquery($1)
LIMIT 10;

-- name: DeleteProject :one
DELETE FROM projects
WHERE external_primary_identifier = $1
RETURNING pk;