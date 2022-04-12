CREATE TABLE IF NOT EXISTS tokens
(
    hash    BYTEA                       NOT NULL PRIMARY KEY,
    user_id BIGINT                      NOT NULL REFERENCES users ON DELETE CASCADE,
    scope   TEXT                        NOT NULL,
    expires TIMESTAMP(0) WITH TIME ZONE NOT NULL
);