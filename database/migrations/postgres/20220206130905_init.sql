-- +goose Up
CREATE TABLE customer
(
    id                 BIGSERIAL PRIMARY KEY,
    email              VARCHAR(128) NOT NULL UNIQUE,
    encrypted_password VARCHAR(128) NOT NULL
);

CREATE TABLE book
(
    id             BIGSERIAL PRIMARY KEY,
    name           VARCHAR(128),
    price          NUMERIC(2) CHECK ( price > 0 ),
    discount_price NUMERIC(2) CHECK ( discount_price > 0 ),
    CONSTRAINT valid_discount CHECK ( price > discount_price )
);

CREATE TABLE author
(
    id         BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(128),
    last_name  VARCHAR(128)
);

CREATE TABLE book_author
(
    book_id   BIGINT REFERENCES book (id) ON DELETE RESTRICT,
    author_id BIGINT REFERENCES author (id) ON DELETE RESTRICT,
    PRIMARY KEY (book_id, author_id)
);

CREATE TABLE warehouse
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(128)
);

CREATE TABLE arrival
(
    id            BIGSERIAL PRIMARY KEY,
    warehouse_id  BIGINT REFERENCES warehouse (id),
    delivery_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE arrival_book
(
    arrival_id BIGINT REFERENCES arrival (id),
    book_id    BIGINT REFERENCES book (id),
    count      int NOT NULL CHECK (count > 0),
    PRIMARY KEY (arrival_id, book_id)
);

CREATE TABLE "order"
(
    id          BIGSERIAL PRIMARY KEY,
    create_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE table order_line
(
    order_id BIGINT REFERENCES "order" (id) ON DELETE CASCADE,
    book_id  BIGINT REFERENCES book (id) ON DELETE RESTRICT,
    count    int NOT NULL CHECK (count > 0),
    PRIMARY KEY (order_id, book_id)
);

-- +goose Down
DROP TABLE IF EXISTS book_author;
DROP TABLE IF EXISTS arrival_book;
DROP TABLE IF EXISTS order_line;
DROP TABLE IF EXISTS customer;
DROP TABLE IF EXISTS book;
DROP TABLE IF EXISTS author;
DROP TABLE IF EXISTS arrival;
DROP TABLE IF EXISTS warehouse;
DROP TABLE IF EXISTS "order";
