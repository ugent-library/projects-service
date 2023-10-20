-- Write your migrate up statements here

ALTER TABLE "projects" ALTER COLUMN "name" TYPE jsonb USING name::jsonb;
ALTER TABLE "projects" ADD COLUMN "ts" tsvector NULL GENERATED ALWAYS AS ((to_tsvector('simple'::regconfig, jsonb_path_query_array(identifier, '$.**{2}'::jsonpath)) || to_tsvector('simple'::regconfig, (id)::text)) || to_tsvector('usimple'::regconfig, jsonb_path_query_array(name, '$.**{2}'::jsonpath))) STORED;

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
