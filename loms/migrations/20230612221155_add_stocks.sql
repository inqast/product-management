-- +goose Up
-- +goose StatementBegin
CREATE TABLE stocks (
    sku bigint NOT NULL,
    warehouse_id bigint NOT NULL,
    count integer NOT NULL,
    reserved integer NOT NULL,
    PRIMARY KEY (sku, warehouse_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
