-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    status integer NOT NULL,
    user_id bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
