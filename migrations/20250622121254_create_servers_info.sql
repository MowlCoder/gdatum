-- +goose Up
-- +goose StatementBegin
CREATE TABLE servers_info
(
    multiplayer LowCardinality(String),
    identifier  String,
    name        String,
    url         String,
    gamemode    String,
    lang        String,
    timestamp   Datetime
) ENGINE = ReplacingMergeTree()
    ORDER BY (identifier, multiplayer)
    PARTITION BY toYYYYMM(timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE servers_info;
-- +goose StatementEnd
