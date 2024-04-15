-- +goose Up
-- +goose StatementBegin
CREATE TABLE payment_status(
    id SERIAL PRIMARY KEY,
    status VARCHAR(25) NOT NULL,
    color VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payment_status;
-- +goose StatementEnd
