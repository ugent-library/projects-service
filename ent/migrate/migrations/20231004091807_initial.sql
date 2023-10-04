-- Create "projects" table
CREATE TABLE "projects" ("id" character varying NOT NULL, "identifier" jsonb NOT NULL, "name" character varying NULL, "description" text NULL, "founding_date" character varying NULL, "dissolution_date" character varying NULL, "acronym" character varying NULL, "grant" character varying NULL, "funding_programme" character varying NULL, "created" timestamptz NOT NULL, "modified" timestamptz NOT NULL, "ts" tsvector NULL GENERATED ALWAYS AS ((to_tsvector('simple'::regconfig, jsonb_path_query_array(identifier, '$.**{2}'::jsonpath)) || to_tsvector('simple'::regconfig, (id)::text)) || to_tsvector('usimple'::regconfig, (name)::text)) STORED, PRIMARY KEY ("id"));
-- Create index "project_ts" to table: "projects"
CREATE INDEX "project_ts" ON "projects" USING gin ("ts");
