CREATE TYPE "Direction" AS ENUM (
  'left',
  'right'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "password" VARCHAR(100) NOT NULL,
  "fullname" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "total_swipes" INT DEFAULT 0,
  "last_swipe_date" TIMESTAMP,
  "swipe_count" INT DEFAULT 0,
  "is_premium" BOOLEAN DEFAULT false
);

CREATE TABLE "swipes" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "target_id" bigint NOT NULL,
  "direction" Direction NOT NULL,
  "swipe_date" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "unique" index_user_target(user_id,target_id)
);

CREATE INDEX ON "swipes" ("user_id", "target_id", "direction");

CREATE INDEX ON "swipes" ("user_id", "swipe_date");

ALTER TABLE "swipes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
