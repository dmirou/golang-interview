-- +goose Up
ALTER TABLE customer
    ADD COLUMN delete_time timestamptz;

-- +goose Down
ALTER TABLE customer
    DROP COLUMN delete_time;