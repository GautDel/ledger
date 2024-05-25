-- +goose Up
-- +goose StatementBegin
ALTER TABLE clients
ADD COLUMN starred BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE clients
DROP COLUMN starred;
-- +goose StatementEnd
