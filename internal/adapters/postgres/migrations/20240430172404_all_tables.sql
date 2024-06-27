-- +goose Up
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
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
    id SERIAL PRIMARY KEY,
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

CREATE TABLE bids (
    id SERIAL PRIMARY KEY,
    item_id BIGINT,
    owner_id BIGINT,
    points FLOAT,

    CONSTRAINT fk_item_id
        FOREIGN KEY (item_id)
            REFERENCES items(id)
                ON DELETE CASCADE,

    CONSTRAINT fk_owner_id
        FOREIGN KEY (owner_id)
            REFERENCES profiles(id)
                ON DELETE CASCADE,

    CONSTRAINT item_points_unique UNIQUE (item_id, points)
);

-- +goose Down
DROP TABLE bids;
DROP TABLE comments;
DROP TABLE items;
DROP TABLE profiles;
