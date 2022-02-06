-- +goose Up
ALTER TABLE customer
    ADD CONSTRAINT not_empty_email CHECK (email != '');

-- +goose Down
ALTER TABLE customer
    DROP CONSTRAINT not_empty_email;