-- create projects

CREATE EXTENSION IF NOT EXISTS "unaccent";
DO
$$BEGIN
  CREATE TEXT SEARCH CONFIGURATION "usimple" ( COPY = simple );
EXCEPTION
  WHEN unique_violation THEN
      NULL;
END;$$;
ALTER TEXT SEARCH CONFIGURATION "usimple" ALTER MAPPING FOR hword, hword_part, word WITH "unaccent", simple;

CREATE TABLE IF NOT EXISTS "projects" (
  "id" character varying NOT NULL,
  "identifier" jsonb NOT NULL,
  "name" character varying NULL,
  "description" text NULL,
  "founding_date" character varying NULL,
  "dissolution_date" character varying NULL,
  "acronym" character varying NULL,
  "grant" character varying NULL,
  "funding_programme" character varying NULL,
  "created" timestamptz NOT NULL,
  "modified" timestamptz NOT NULL,
  "ts" tsvector NULL GENERATED ALWAYS AS ((to_tsvector('simple'::regconfig, jsonb_path_query_array(identifier, '$.**{2}'::jsonpath)) || to_tsvector('simple'::regconfig, (id)::text)) || to_tsvector('usimple'::regconfig, (name)::text)) STORED,
  PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS "project_ts" ON "projects" USING gin ("ts");

---- create above / drop below ----

DROP TABLE "projects";
DROP INDEX "projects_ts";