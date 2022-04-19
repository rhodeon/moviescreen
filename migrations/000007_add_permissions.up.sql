CREATE TABLE IF NOT EXISTS permissions
(
    id   BIGSERIAL NOT NULL PRIMARY KEY,
    code TEXT      NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions
(
    user_id       BIGINT NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id BIGINT NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);

-- add permissions
INSERT INTO permissions(code)
VALUES ('movies:read'),
       ('movies:write');

