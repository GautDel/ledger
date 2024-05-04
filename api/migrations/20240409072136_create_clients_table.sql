-- +goose Up
-- +goose StatementBegin
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    description text,
    email VARCHAR(320),
    phone VARCHAR(50),
    address VARCHAR(1000),
    country VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;
-- +goose StatementEnd

