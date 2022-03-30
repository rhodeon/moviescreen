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

INSERT INTO movies(title, year, runtime, genres)
VALUES ('Bullet Train', 2022, 108, '{"Action", "Comedy"}');

INSERT INTO movies(title, year, runtime, genres)
VALUES ('Hamilton', 2020, 140, '{"Musical", "Drama"}');

INSERT INTO movies(title, year, runtime, genres)
VALUES ('Luca', 2021, 100, '{"Adventure", "Family"}');
