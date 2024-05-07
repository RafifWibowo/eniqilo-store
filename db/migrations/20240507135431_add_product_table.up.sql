CREATE TYPE ProdCategories AS ENUM (
    'Clothing',
    'Accessories',
    'Footwear',
    'Beverages'
);

CREATE TABLE IF NOT EXISTS "product" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "name" varchar(30) NOT NULL,
    "sku" varchar(30) NOT NULL,
    "category" ProdCategories NOT NULL,
    "imageUrl" varchar(255) NOT NULL,
    "notes" varchar(200) NOT NULL,
    "price" int NOT NULL,
    "stock" int NOT NULL,
    "location" varchar(200) NOT NULL,
    "isAvailable" bool NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now()),
    "deleted_at" timestamp
);