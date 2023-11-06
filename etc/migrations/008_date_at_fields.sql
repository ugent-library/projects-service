-- Write your migrate up statements here

ALTER TABLE "projects" ADD COLUMN "created_at" timestamptz;
UPDATE projects SET "created" = "created_at";
ALTER TABLE "projects" ALTER COLUMN "created_at" SET NOT NULL;

ALTER TABLE "projects" ADD COLUMN "updated_at" timestamptz;
UPDATE projects SET "modified" = "updated_at";
ALTER TABLE "projects" ALTER COLUMN "updated_at" SET NOT NULL;

ALTER TABLE "projects" DROP COLUMN "created";
ALTER TABLE "projects" DROP COLUMN "modified";

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
