-- +goose Up
-- +goose StatementBegin
ALTER TABLE invoices
ADD COLUMN status VARCHAR(50) DEFAULT 'unpaid';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE invoices
DROP COLUMN status;
-- +goose StatementEnd
