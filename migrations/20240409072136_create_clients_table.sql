-- +goose Up
-- +goose StatementBegin
CREATE TABLE clients (
    id int NOT NULL,
    first_name text,
    last_name text,
    user_id text NOT NULL UNIQUE,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;
-- +goose StatementEnd
