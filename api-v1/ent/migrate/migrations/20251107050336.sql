-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "deleted" boolean NOT NULL DEFAULT false;
