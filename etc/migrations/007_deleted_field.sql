-- Write your migrate up statements here

ALTER TABLE "projects" ADD COLUMN "deleted" BOOLEAN;
UPDATE "projects" SET "deleted" = 'false';
ALTER TABLE "projects" ALTER COLUMN "deleted" SET NOT NULL;
ALTER TABLE "projects" ALTER COLUMN "deleted" SET DEFAULT FALSE;

---- create above / drop below ----

ALTER TABLE "projects" DROP COLUMN "deleted";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.