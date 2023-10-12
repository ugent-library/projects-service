-- Write your migrate up statements here

ALTER TABLE "projects" DROP COLUMN "grant", ADD COLUMN "grant_id" character varying NULL;

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
