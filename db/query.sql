-- name: GetProject :one
SELECT p.* 
FROM projects p, projects_identifiers pi
WHERE p.id = pi.project_id AND pi.type = $1 AND pi.value = $2;

-- name: CreateProject :one
INSERT INTO projects(
    name,
    description,
    founding_date,
    dissolution_date,
    attributes
)
VAlUES($1, $2, $3, $4, $5)
RETURNING id;

-- name: UpdateProject :exec
UPDATE projects SET (
    name,
    description,
    founding_date,
    dissolution_date,
    attributes,
    updated_at
) = ($2, $3, $4, $5, $6, CURRENT_TIMESTAMP)
WHERE id = $1;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1;

-- name: GetProjectIdentifiers :many
SELECT * FROM projects_identifiers
WHERE project_id = $1;

-- name: CreateProjectIdentifier :exec
INSERT INTO projects_identifiers(
    project_id,
    type,
    value
) VALUES ($1, $2, $3);

-- name: DeleteProjectIdentifier :exec
DELETE FROM projects_identifiers
WHERE type = $1 AND value = $2;

-- name: EachProject :many
SELECT 
    name,
    description,
    founding_date,
    dissolution_date,
    created_at,
    updated_at
FROM projects
ORDER BY id ASC 
OFFSET $1
LIMIT $2;

-- name: BetweenProjects :many
SELECT
    name,
    description,
    founding_date,
    dissolution_date,
    created_at,
    updated_at
FROM projects
WHERE created_at >= $1 AND created_at <= $2
ORDER BY created_at ASC;