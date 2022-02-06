-- +goose Up
ALTER TABLE customer
    ALTER first_name SET NOT NULL,
    ALTER first_name SET DEFAULT '',
    ALTER last_name SET NOT NULL,
    ALTER last_name SET DEFAULT '';

-- +goose Down
ALTER TABLE customer
    ALTER first_name DROP NOT NULL,
    ALTER first_name DROP DEFAULT,
    ALTER last_name DROP NOT NULL,
    ALTER last_name DROP DEFAULT;
