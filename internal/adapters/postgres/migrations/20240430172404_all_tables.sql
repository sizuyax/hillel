-- +goose Up
-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     email VARCHAR(50) UNIQUE,
--     password VARCHAR(256)
-- );

-- CREATE TABLE sellers(
--     id SERIAL PRIMARY KEY ,
--     email VARCHAR(50) UNIQUE,
--     password VARCHAR(256)
-- );

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY ,
    email VARCHAR(50) UNIQUE,
    password VARCHAR(256),
    type VARCHAR(20)
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    owner_id BIGINT,
    name VARCHAR(256) UNIQUE,
    price FLOAT,

    CONSTRAINT fk_owner_id
        FOREIGN KEY (owner_id)
            REFERENCES profiles(id)
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY ,
    item_id BIGINT,
    owner_id BIGINT,
    body TEXT UNIQUE,

    CONSTRAINT fk_item_id
        FOREIGN KEY (item_id)
            REFERENCES items(id)
                      ON DELETE CASCADE,

    CONSTRAINT fk_owner_id
        FOREIGN KEY (owner_id)
            REFERENCES profiles(id)
                      ON DELETE CASCADE
);

-- +goose Down
DROP TABLE comments;
DROP TABLE items;
DROP TABLE profiles;