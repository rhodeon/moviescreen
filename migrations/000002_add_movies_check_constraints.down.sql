ALTER TABLE IF EXISTS movies
    DROP CONSTRAINT IF EXISTS movies_runtime_check;

ALTER TABLE IF EXISTS movies
    DROP CONSTRAINT IF EXISTS movies_year_check;

ALTER TABLE IF EXISTS movies
    DROP CONSTRAINT IF EXISTS genres_length_check;