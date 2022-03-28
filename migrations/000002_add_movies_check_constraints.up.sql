ALTER TABLE movies
    ADD CONSTRAINT movies_runtime_check CHECK ( runtime > 0 );

ALTER TABLE movies
    ADD CONSTRAINT movies_year_check CHECK ( year BETWEEN 1888 and DATE_PART('year', NOW()));

ALTER TABLE movies
    ADD CONSTRAINT genres_length_check CHECK (ARRAY_LENGTH(genres, 1) > 0);