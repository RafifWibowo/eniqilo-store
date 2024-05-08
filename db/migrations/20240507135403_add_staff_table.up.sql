CREATE TABLE IF NOT EXISTS "staff" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "phoneNumber" varchar(20) UNIQUE NOT NULL,
    "password" varchar(255) NOT NULL,
    "createdAt" timestamp NOT NULL DEFAULT (now()),
    "updatedAt" timestamp NOT NULL DEFAULT (now()),
    "deletedAt" timestamp
);