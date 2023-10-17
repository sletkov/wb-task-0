-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    order_uid varchar not null,
    track_number varchar not null,
    entry varchar not null,
    delivery json not null,
    payment json not null,
    items json not null,
    locale varchar not null,
    internal_signature varchar,
    customer_id varchar not null,
    delivery_service varchar not null,
    shardkey varchar not null,
    sm_id integer not null,
    date_created varchar not null,
    oof_shard varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
