-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
    order_id bigint NOT NULL,
    sku bigint NOT NULL,
    count integer NOT NULL,
    PRIMARY KEY (order_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
