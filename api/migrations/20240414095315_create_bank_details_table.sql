-- +goose Up
-- +goose StatementBegin
CREATE TABLE bank_details(
    id UUID PRIMARY KEY,
    bic BYTEA NOT NULL,
    iban BYTEA NOT NULL,
    account_name BYTEA NOT NULL,
    bank_name BYTEA NOT NULL,
    bank_location BYTEA NOT NULL,
    user_id VARCHAR(255) REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bank_details;
-- +goose StatementEnd
