-- +goose Up
-- +goose StatementBegin
CREATE TABLE invoice_items(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invoice_id UUID REFERENCES invoices(id), 
    qty SMALLINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    unit_price NUMERIC(10, 2),
    hourly_price NUMERIC(10, 2),
    total_price NUMERIC(10, 2),
    user_id VARCHAR(255) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoice_items;
-- +goose StatementEnd
