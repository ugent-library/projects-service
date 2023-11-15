-- create projects
CREATE EXTENSION IF NOT EXISTS "unaccent";
DO $$BEGIN CREATE TEXT SEARCH CONFIGURATION "usimple" (COPY = simple);
EXCEPTION
WHEN unique_violation THEN NULL;
END;
$$;
ALTER TEXT SEARCH CONFIGURATION "usimple" ALTER MAPPING FOR hword,
hword_part,
word WITH "unaccent",
simple;

CREATE TABLE IF NOT EXISTS "projects" (
  "pk" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY,
  "external_primary_identifier" text NOT NULL,
  "external_identifiers" jsonb NOT NULL,
  "name" jsonb NOT NULL,
  "description" jsonb NULL,
  "founding_date" text NULL,
  "dissolution_date" text NULL,
  "acronym" jsonb NULL,
  "eu_grant_call" text NULL,
  "eu_funding_programme" text NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "ts" tsvector NULL GENERATED ALWAYS AS (
    to_tsvector(
      'simple'::regconfig,
      jsonb_path_query_array(external_identifiers, '$.**{2}'::jsonpath)
    ) || to_tsvector(
      'usimple'::regconfig,
      jsonb_path_query_array(name, '$.**{2}'::jsonpath)
    )
  ) STORED,
  PRIMARY KEY ("pk"),
  UNIQUE ("external_primary_identifier")
);

CREATE INDEX IF NOT EXISTS "project_ts_idx" ON "projects" USING gin ("ts");

---- create above / drop below ----
DROP TABLE "projects";
DROP INDEX "project_ts_idx";