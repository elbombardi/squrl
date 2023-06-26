CREATE TABLE "customers" (
  "id" int PRIMARY KEY,
  "prefix" varchar(3) UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "api_key" varchar UNIQUE NOT NULL,
  "status" varchar(1) NOT NULL DEFAULT 'e',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "short_urls" (
  "id" int PRIMARY KEY,
  "short_url_key" varchar,
  "customer_id" int NOT NULL,
  "long_url" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'e',
  "tracking_status" varchar(1) DEFAULT 'e',
  "click_count" int DEFAULT 0,
  "first_click_date_time" timestamp,
  "last_click_date_time" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "clicks" (
  "id" int PRIMARY KEY,
  "short_url_id" int NOT NULL,
  "click_date_time" timestamp DEFAULT (now()),
  "user_agent" varchar,
  "ip_address" varchar
);

CREATE INDEX ON "customers" ("prefix");

CREATE INDEX ON "customers" ("username");

CREATE INDEX ON "customers" ("api_key");

CREATE INDEX ON "short_urls" ("customer_id");

CREATE UNIQUE INDEX ON "short_urls" ("short_url_key", "customer_id");

CREATE INDEX ON "clicks" ("short_url_id");

COMMENT ON TABLE "customers" IS 'Table holding Customer information';

COMMENT ON COLUMN "customers"."prefix" IS '3 characters, case-sensitive';

COMMENT ON COLUMN "customers"."api_key" IS 'API key';

COMMENT ON COLUMN "customers"."status" IS 'e: enabled, d: disabled';

COMMENT ON COLUMN "customers"."created_at" IS 'Timestamp of creation';

COMMENT ON COLUMN "customers"."updated_at" IS 'Timestamp of last update';

COMMENT ON TABLE "short_urls" IS 'Table holding short URL information';

COMMENT ON COLUMN "short_urls"."short_url_key" IS '6 characters, case-sensitive';

COMMENT ON COLUMN "short_urls"."status" IS 'e: enabled, d: disabled';

COMMENT ON COLUMN "short_urls"."tracking_status" IS 'e: enabled, d: disabled';

COMMENT ON COLUMN "short_urls"."click_count" IS 'Aggregate updated by the redirection server';

COMMENT ON COLUMN "short_urls"."first_click_date_time" IS 'Aggregate set by the redirection server';

COMMENT ON COLUMN "short_urls"."last_click_date_time" IS 'Aggregate set by the redirection server';

COMMENT ON COLUMN "short_urls"."created_at" IS 'Timestamp of creation';

COMMENT ON COLUMN "short_urls"."updated_at" IS 'Timestamp of last update';

COMMENT ON TABLE "clicks" IS 'Table holding click information';

COMMENT ON COLUMN "clicks"."click_date_time" IS 'Timestamp of click';

ALTER TABLE "short_urls" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "clicks" ADD FOREIGN KEY ("short_url_id") REFERENCES "short_urls" ("id");
