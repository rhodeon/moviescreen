ALTER TABLE IF EXISTS users
    ADD CONSTRAINT users_username_length_check CHECK ( length(username) <= 500);