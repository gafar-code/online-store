CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "address" varchar NOT NULL,
  "token" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "virtual_accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "rekening_number" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "category_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" bigint NOT NULL,
  "qty" bigint NOT NULL DEFAULT '1',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "qty" bigint NOT NULL DEFAULT '1',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "status" varchar NOT NULL,
  "virtual_account_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "qty" bigint NOT NULL DEFAULT '1',
  "order_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'PENDING',
  "issued_at" timestamptz NOT NULL DEFAULT (now()),
  "order_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "customers" ("id");

CREATE INDEX ON "virtual_accounts" ("id");

CREATE INDEX ON "product_categories" ("id");

CREATE INDEX ON "products" ("id");

CREATE INDEX ON "carts" ("id");

CREATE INDEX ON "orders" ("id");

CREATE INDEX ON "transactions" ("id");

COMMENT ON COLUMN "orders"."status" IS 'WAITING_PAYMENT/PAID/CANCEL';

COMMENT ON COLUMN "transactions"."status" IS 'IN_DELIVERY/DELIVERED/SUCCESS';

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "product_categories" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("virtual_account_id") REFERENCES "virtual_accounts" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
