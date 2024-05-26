-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    email VARCHAR(50) UNIQUE,
    password VARCHAR(256)
);

CREATE TABLE sellers(
    id SERIAL PRIMARY KEY ,
    email VARCHAR(50) UNIQUE,
    password VARCHAR(256)
);

CREATE TABLE items (
    id SERIAL,
    owner_id BIGINT,
    name VARCHAR(256),
    price FLOAT,

    CONSTRAINT fk_owner_id
        FOREIGN KEY (owner_id)
            REFERENCES sellers(id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE sellers;
DROP TABLE items;