-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks (sku, warehouse_id, count, reserved)
VALUES
    (1076963, 1, 3, 0),
    (1076963, 2, 3, 0),
    (1076963, 3, 3, 0),
    (1148162, 1, 3, 0),
    (1148162, 2, 3, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
