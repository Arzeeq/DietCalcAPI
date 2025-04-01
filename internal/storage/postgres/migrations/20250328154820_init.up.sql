CREATE TABLE IF NOT EXISTS "users" (
    "login" VARCHAR(20),
    "password" VARCHAR NOT NULL,
    "sex" VARCHAR,
    "birthdate" DATE,
    "height" SMALLINT,
    "purpose" VARCHAR,
    "created_at" TIMESTAMP NOT NULL,
    CONSTRAINT "PK_USERS_LOGIN" PRIMARY KEY ("login")
);

CREATE TABLE IF NOT EXISTS "products" (
    "id" SERIAL,
    "name" VARCHAR,
    "calories" NUMERIC(5, 1),
    "proteins" NUMERIC(5, 1),
    "fats" NUMERIC(5, 1),
    "carbs" NUMERIC(5, 1),
    "user_login" VARCHAR(20),
    CONSTRAINT "PK_PRODUCTS_ID" PRIMARY KEY ("id"),
    CONSTRAINT "FK_PRODUCTS_USERLOGIN" FOREIGN KEY("user_login") REFERENCES "users"("login") ON DELETE SET NULL
);