-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hello_world_entity (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS hello_world_entity;
-- +goose StatementEnd
