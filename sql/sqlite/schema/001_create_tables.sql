-- +goose Up
CREATE TABLE IF NOT EXISTS closs_session(
    refresh_token TEXT NOT NULL,
    access_token TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    user_id TEXT NOT NULL PRIMARY KEY
);

-- +goose Down
DROP TABLE closs_session;

