-- +goose Up
-- +goose StatementBegin
CREATE TABLE invoices(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invoice_id VARCHAR(100) NOT NULL,
    comp_name VARCHAR(255) NOT NULL,
    comp_addr VARCHAR(255),
    comp_email VARCHAR(255),
    comp_phone VARCHAR(255),
    sub_total numeric(12, 2),
    bank_acc_name BYTEA, 
    bic BYTEA, 
    iban BYTEA, 
    invoice_date DATE DEFAULT CURRENT_DATE,
    due_date DATE,
    client_name VARCHAR(255) NOT NULL,
    client_addr VARCHAR(255),
    client_phone VARCHAR(255),
    client_email VARCHAR(255),
    tax_percent SMALLINT,
    user_id VARCHAR(255) REFERENCES users(id),
    client_id INT references clients(id),
    project_id INT references projects(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoices;
-- +goose StatementEnd
