-- movies
INSERT INTO movies(title, year, runtime, genres)
VALUES ('Bullet Train', 2022, 108, '{"Action", "Comedy"}'),
       ('Hamilton', 2020, 140, '{"Musical", "Drama"}'),
       ('Luca', 2021, 100, '{"Adventure", "Family"}');

-- users
INSERT INTO users(created_at, username, email, password_hash, activated, version)
VALUES (now(), 'rhodeon', 'rhodeon@dev.mail', '$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm',
        true, 1),
       (now(), 'ruona', 'ruona@mail.com', '$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm',
        false, 1),
       (now(), 'johndoe', 'johndoe@mail.com', '$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm',
        true, 3);

-- tokens
INSERT INTO tokens
VALUES ('\x33a70a01f980448b333d89fecae178754529ef424eca7794298d658f257a5959', 1, 'activation', now() + interval '2h'),
       ('\x1937df93f58c4aa1baad9517ccea28fd391bf7a5e85c35ef70c211cbaa124bb7', 1, 'authentication',
        '1970-02-01 00:00:00-00');

-- permissions
INSERT INTO permissions(code)
VALUES ('movies:read'),
       ('movies:write'),
       ('metrics:view');

-- users_permissions
INSERT INTO users_permissions(user_id, permission_id)
VALUES (1, 1),
       (1, 2);