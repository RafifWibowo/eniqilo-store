CREATE TABLE IF NOT EXISTS "customer" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "phone" varchar(20) UNIQUE NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now()),
    "deleted_at" timestamp
);