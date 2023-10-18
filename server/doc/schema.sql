-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-10-18T21:59:35.916Z

CREATE TABLE "users" (
  "username" varchar,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "is_user_verified" bool NOT NULL DEFAULT false,
  PRIMARY KEY ("username")
);

CREATE TABLE "admin_user" (
  "admin_key" varchar
);
