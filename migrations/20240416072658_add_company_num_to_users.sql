-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN company_num VARCHAR(255) DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN company_num;
-- +goose StatementEnd
