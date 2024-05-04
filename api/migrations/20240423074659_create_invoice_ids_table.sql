-- +goose Up
-- +goose StatementBegin
CREATE TABLE invoice_ids (
    id VARCHAR(10) PRIMARY KEY NOT NULL,
    invoice_id UUID REFERENCES invoices(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoice_ids;
-- +goose StatementEnd
