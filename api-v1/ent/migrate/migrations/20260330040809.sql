-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "supabase" character varying NULL;
-- Create index "users_supabase_key" to table: "users"
CREATE UNIQUE INDEX "users_supabase_key" ON "users" ("supabase");
