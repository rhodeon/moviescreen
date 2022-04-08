CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    created_at    TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    username      CITEXT                      NOT NULL UNIQUE,
    email         CITEXT                      NOT NULL UNIQUE,
    password_hash BYTEA                       NOT NULL,
    activated     BOOL                        NOT NULL DEFAULT FALSE,
    version       INTEGER                     NOT NULL DEFAULT 1
);