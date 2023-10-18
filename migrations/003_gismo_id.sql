-- Write your migrate up statements here

ALTER TABLE "projects" ADD COLUMN "gismo_id" character varying NOT NULL;

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
