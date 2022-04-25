-- movies table
DROP TABLE IF EXISTS movies;
CREATE TABLE IF NOT EXISTS movies
(
    id         bigserial PRIMARY KEY       NOT NULL,
    title      text                        NOT NULL,
    year       integer                     NOT NULL,
    runtime    integer                     NOT NULL,
    genres     text[]                      NOT NULL,
    version    integer                     NOT NULL DEFAULT 1,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

ALTER TABLE movies
    ADD CONSTRAINT movies_runtime_check CHECK ( runtime > 0 );

ALTER TABLE movies
    ADD CONSTRAINT movies_year_check CHECK ( year BETWEEN 1888 and DATE_PART('year', NOW()));

ALTER TABLE movies
    ADD CONSTRAINT genres_length_check CHECK (ARRAY_LENGTH(genres, 1) > 0);

CREATE INDEX IF NOT EXISTS movies_title_idx ON movies USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS movies_genres_idx ON movies USING GIN (genres);

INSERT INTO movies(title, year, runtime, genres)
VALUES ('Bullet Train', 2022, 108, '{"Action", "Comedy"}');

INSERT INTO movies(title, year, runtime, genres)
VALUES ('Hamilton', 2020, 140, '{"Musical", "Drama"}');

INSERT INTO movies(title, year, runtime, genres)
VALUES ('Luca', 2021, 100, '{"Adventure", "Family"}');

-- users table
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

ALTER TABLE users
    ADD CONSTRAINT users_username_length_check CHECK ( length(username) <= 500);

INSERT INTO users(created_at, username, email, password_hash, activated, version)
VALUES (now(), 'rhodeon', 'rhodeon@dev.mail', '$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm',
        true, 1),
       (now(), 'ruona', 'ruona@mail.com', '$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm',
        false, 1),
       (now(), 'johndoe', 'johndoe@mail.com', '$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm',
        true, 3);


-- tokens table
CREATE TABLE IF NOT EXISTS tokens
(
    hash    BYTEA                       NOT NULL PRIMARY KEY,
    user_id BIGINT                      NOT NULL REFERENCES users ON DELETE CASCADE,
    scope   TEXT                        NOT NULL,
    expires TIMESTAMP(0) WITH TIME ZONE NOT NULL
);

-- permissions table
CREATE TABLE IF NOT EXISTS permissions
(
    id   BIGSERIAL NOT NULL PRIMARY KEY,
    code TEXT      NOT NULL
);

INSERT INTO permissions(code)
VALUES ('movies:read'),
       ('movies:write'),
       ('metrics:view');

-- users_permissions table
CREATE TABLE IF NOT EXISTS users_permissions
(
    user_id       BIGINT NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id BIGINT NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);