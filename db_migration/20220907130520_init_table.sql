-- +goose Up
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS data(
    id text not null,
    data json
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE data;