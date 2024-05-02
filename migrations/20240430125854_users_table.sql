-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    email VARCHAR(50) UNIQUE,
    password VARCHAR(256)
);


-- +goose Down
DROP TABLE users;

