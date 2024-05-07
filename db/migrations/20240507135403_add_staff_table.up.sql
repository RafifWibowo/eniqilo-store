CREATE TABLE IF NOT EXISTS "staff" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "phone" varchar(20) UNIQUE NOT NULL,
    "password" varchar(255) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now()),
    "deleted_at" timestamp
);