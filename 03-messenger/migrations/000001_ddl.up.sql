CREATE TABLE "users" (
  "user_id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "messages" (
  "message_id" SERIAL PRIMARY KEY,
  "sender_id" int NOT NULL,
  "receiver_id" int NOT NULL,
  "subject" varchar NOT NULL DEFAULT '',
  "body" text NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

COMMENT ON TABLE "messages" IS 'Unique Constraint (sender_id, receiver_id, message_id)';

ALTER TABLE "messages" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("user_id");

ALTER TABLE "messages" ADD FOREIGN KEY ("receiver_id") REFERENCES "users" ("user_id");
