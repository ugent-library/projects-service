-- Write your migrate up statements here

CREATE UNIQUE INDEX "projects_gismo_id_key" ON "projects" ("gismo_id");

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
