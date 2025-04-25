CREATE TYPE "register_status" AS ENUM (
  'created',
  'running',
  'done',
  'evaluated',
  'failure'
);

CREATE TABLE "risks" (
  "uid" uuid PRIMARY KEY,
  "details" string,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "users" (
  "uid" uuid PRIMARY KEY,
  "username" varchar,
  "role" varchar,
  "created_at" timestamp DEFAULT (now()),
  "deleted_at" bool DEFAULT 'false'
);

CREATE TABLE "user_risk" (
  "user_risk_uid" uuid PRIMARY KEY,
  "user_uid" uuid,
  "risk_uid" uuid,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "images" (
  "uid" uuid PRIMARY KEY,
  "bucket_name" varchar,
  "bucket_id" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "audios" (
  "uid" uuid PRIMARY KEY,
  "bucket_name" varchar,
  "bucket_id" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "registers" (
  "uid" uuid PRIMARY KEY,
  "title" varchar,
  "body" text,
  "risk_scale" integer NOT NULL,
  "local" varchar,
  "status" register_status,
  "image_uid" uuid,
  "audios_uid" uuid,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_registers" (
  "uid" uuid PRIMARY KEY,
  "user_uid" uuid,
  "register_uid" uuid,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "configuration" (
  "uid" uuid PRIMARY KEY,
  "config" jsonb,
  "user_uid" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "preferences" (
  "uid" uuid PRIMARY KEY,
  "prefer" jsonb,
  "user_uid" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

COMMENT ON COLUMN "registers"."body" IS 'Content of the post';

ALTER TABLE "user_risk" ADD FOREIGN KEY ("user_uid") REFERENCES "users" ("uid");

ALTER TABLE "configuration" ADD FOREIGN KEY ("user_uid") REFERENCES "users" ("uid");

ALTER TABLE "preferences" ADD FOREIGN KEY ("user_uid") REFERENCES "users" ("uid");

ALTER TABLE "user_registers" ADD FOREIGN KEY ("user_uid") REFERENCES "users" ("uid");

ALTER TABLE "user_risk" ADD FOREIGN KEY ("risk_uid") REFERENCES "risks" ("uid");

ALTER TABLE "user_registers" ADD FOREIGN KEY ("register_uid") REFERENCES "registers" ("uid");

ALTER TABLE "registers" ADD FOREIGN KEY ("image_uid") REFERENCES "images" ("uid");

ALTER TABLE "registers" ADD FOREIGN KEY ("audios_uid") REFERENCES "audios" ("uid");
