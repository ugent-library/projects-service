-- +goose Up

CREATE TABLE IF NOT EXISTS projects (
  id BIGSERIAL PRIMARY KEY,
  name JSONB NULL,
  description JSONB NULL,
  founding_date TEXT NULL,
  dissolution_date TEXT NULL,
  attributes JSONB,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX projects_updated_at_key ON projects (updated_at);

CREATE TABLE IF NOT EXISTS "projects_identifiers" (
  project_id BIGINT NOT NULL REFERENCES projects ON DELETE CASCADE,
  type TEXT CHECK (type <> ''),
  value TEXT CHECK (type <> ''),
  PRIMARY KEY (type, value)
);

CREATE INDEX projects_identifiers_project_id_fkey ON projects_identifiers (project_id);

-- +goose Down

DROP TABLE "projects" CASCADE;
DROP TABLE "projects_identifiers" CASCADE;