-- +goose Up
ALTER TABLE customer
    DROP COLUMN encrypted_password;

-- +goose Down
ALTER TABLE customer
    ADD COLUMN encrypted_password VARCHAR(128) NOT NULL;
