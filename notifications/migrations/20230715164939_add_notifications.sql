-- +goose Up
-- +goose StatementBegin
CREATE TABLE notifications (
   user_id bigint NOT NULL,
   order_id bigint NOT NULL,
   status varchar NOT NULL,
   created_at  timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications;
-- +goose StatementEnd
