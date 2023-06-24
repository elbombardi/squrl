CREATE TABLE "customer" (
  "id" int PRIMARY KEY,
  "prefix" varchar(3) UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "api_key" varchar UNIQUE NOT NULL,
  "status" varchar(1) NOT NULL DEFAULT 'e',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "short_url" (
  "id" int PRIMARY KEY,
  "short_url_key" varchar,
  "customer_id" int NOT NULL,
  "long_url" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'e',
  "click_count" int DEFAULT 0,
  "first_click_date_time" timestamp,
  "last_click_date_time" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "click" (
  "id" int PRIMARY KEY,
  "short_url_id" int NOT NULL,
  "click_date_time" timestamp DEFAULT (now()),
  "user_agent" varchar,
  "ip_address" varchar
);

CREATE INDEX ON "customer" ("prefix");

CREATE INDEX ON "customer" ("username");

CREATE INDEX ON "customer" ("api_key");

CREATE INDEX ON "short_url" ("customer_id");

CREATE UNIQUE INDEX ON "short_url" ("short_url_key", "customer_id");

CREATE INDEX ON "click" ("short_url_id");

COMMENT ON TABLE "customer" IS 'Table holding Customer information';

COMMENT ON COLUMN "customer"."prefix" IS '3 characters, case-sensitive';

COMMENT ON COLUMN "customer"."api_key" IS 'API key';

COMMENT ON COLUMN "customer"."status" IS 'e: enabled, d: disabled';

COMMENT ON COLUMN "customer"."created_at" IS 'Timestamp of creation';

COMMENT ON COLUMN "customer"."updated_at" IS 'Timestamp of last update';

COMMENT ON TABLE "short_url" IS 'Table holding short URL information';

COMMENT ON COLUMN "short_url"."short_url_key" IS '6 characters, case-sensitive';

COMMENT ON COLUMN "short_url"."status" IS 'e: enabled, d: disabled';

COMMENT ON COLUMN "short_url"."click_count" IS 'Aggregate updated by the redirection server';

COMMENT ON COLUMN "short_url"."first_click_date_time" IS 'Aggregate set by the redirection server';

COMMENT ON COLUMN "short_url"."last_click_date_time" IS 'Aggregate set by the redirection server';

COMMENT ON COLUMN "short_url"."created_at" IS 'Timestamp of creation';

COMMENT ON COLUMN "short_url"."updated_at" IS 'Timestamp of last update';

COMMENT ON TABLE "click" IS 'Table holding click information';

COMMENT ON COLUMN "click"."click_date_time" IS 'Timestamp of click';

ALTER TABLE "customer" ADD FOREIGN KEY ("id") REFERENCES "short_url" ("customer_id");

ALTER TABLE "short_url" ADD FOREIGN KEY ("id") REFERENCES "click" ("short_url_id");
