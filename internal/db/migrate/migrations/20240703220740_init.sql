-- Create "tokens" table
CREATE TABLE "tokens" ("id" uuid NOT NULL, "access_token" character varying NOT NULL, "refresh_token" character varying NOT NULL, "expire_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "username" character varying NOT NULL, "email" character varying NOT NULL, "external_user_id" bigint NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "user_tokens" table
CREATE TABLE "user_tokens" ("user_id" uuid NOT NULL, "token_id" uuid NOT NULL, PRIMARY KEY ("user_id", "token_id"), CONSTRAINT "user_tokens_token_id" FOREIGN KEY ("token_id") REFERENCES "tokens" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_tokens_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
