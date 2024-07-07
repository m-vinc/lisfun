-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "username" DROP NOT NULL, ADD COLUMN "first_name" character varying NULL;
