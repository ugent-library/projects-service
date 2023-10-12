-- Modify "projects" table
ALTER TABLE "projects" DROP COLUMN "grant", ADD COLUMN "grant_id" character varying NULL;
