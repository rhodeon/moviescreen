ALTER TABLE IF EXISTS movies
    ADD CONSTRAINT movies_runtime_check CHECK ( runtime > 0 );

ALTER TABLE IF EXISTS movies
    ADD CONSTRAINT movies_year_check CHECK ( year BETWEEN 1888 and date_part('year', now()));

ALTER TABLE IF EXISTS movies
    ADD CONSTRAINT genres_length_check CHECK (array_length(genres, 1) > 0);