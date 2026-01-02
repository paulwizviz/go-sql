-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL
);

INSERT INTO users VALUES
(0, 'root', '', ''),
(1, 'vojtechvitek', 'Vojtech', 'Vitek');

-- +goose Down
DROP TABLE users;