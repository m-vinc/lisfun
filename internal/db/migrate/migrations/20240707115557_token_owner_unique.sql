-- Modify "tokens" table
ALTER TABLE "tokens" ADD COLUMN "user_tokens" uuid NULL, ADD CONSTRAINT "tokens_users_tokens" FOREIGN KEY ("user_tokens") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Drop "user_tokens" table
DROP TABLE "user_tokens";
