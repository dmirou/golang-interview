-- +goose Up
ALTER TABLE customer
    ADD COLUMN first_name VARCHAR(128) NULL,
    ADD COLUMN last_name VARCHAR(128) NULL;

-- +goose Down
ALTER TABLE customer
    DROP COLUMN first_name,
    DROP COLUMN last_name;
