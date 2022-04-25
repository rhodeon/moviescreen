-- users_permissions
DROP TABLE IF EXISTS users_permissions;

-- permissions
DROP TABLE IF EXISTS permissions;

-- tokens
DROP TABLE IF EXISTS tokens;

-- users
ALTER TABLE users
    DROP CONSTRAINT IF EXISTS users_username_length_check;

DROP TABLE IF EXISTS users;

-- movies
DROP TABLE IF EXISTS movies;
