-- create projects

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

CREATE TABLE IF NOT EXISTS "projects_identifiers" (
  project_id BIGINT NOT NULL REFERENCES projects ON DELETE CASCADE,
  type TEXT CHECK (type <> ''),
  value TEXT CHECK (type <> ''),
  PRIMARY KEY (type, value)
);

---- create above / drop below ----
DROP TABLE "projects" CASCADE;
DROP TABLE "projects_identifiers" CASCADE;