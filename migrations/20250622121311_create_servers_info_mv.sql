-- +goose Up
-- +goose StatementBegin
CREATE
MATERIALIZED VIEW servers_info_mv TO servers_info AS
SELECT multiplayer,
       identifier,
       name,
       url,
       gamemode,
       lang,
       timestamp
FROM servers_metrics_raw
GROUP BY multiplayer, identifier, name, url, gamemode, lang, timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE servers_info_mv;
-- +goose StatementEnd
