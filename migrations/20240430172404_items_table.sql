-- +goose Up
CREATE TABLE items (
    id SERIAL,
    owner_id int,
    name VARCHAR(256),
    price float,

    CONSTRAINT fk_owner_id
        FOREIGN KEY (owner_id)
            REFERENCES users(id)
);

-- +goose Down
DROP TABLE items;
