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
SELECT p.*, json_agg(json_build_object('type', pi.type, 'value', pi.value)) AS identifiers
FROM projects p
LEFT JOIN projects_identifiers pi ON p.id = pi.project_id
WHERE p.id > $1
GROUP BY p.id LIMIT $2;