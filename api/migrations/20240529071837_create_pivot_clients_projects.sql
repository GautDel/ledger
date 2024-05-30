-- +goose Up
-- +goose StatementBegin
CREATE TABLE clients_projects(
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id),
    project_id INT REFERENCES projects(id),
    user_id VARCHAR(255) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients_projects;
-- +goose StatementEnd
