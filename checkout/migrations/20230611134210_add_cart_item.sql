-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
    user_id bigint NOT NULL,
    sku bigint NOT NULL,
    count integer NOT NULL,
    PRIMARY KEY (user_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd

