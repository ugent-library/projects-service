-- Write your migrate up statements here

ALTER TABLE "projects" ALTER COLUMN "description" TYPE jsonb USING name::jsonb;

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
