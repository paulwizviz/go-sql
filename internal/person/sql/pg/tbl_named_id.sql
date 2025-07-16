CREATE TABLE IF NOT EXISTS named_identifier (
    id SERIAL PRIMARY KEY,
    -- VARCHAR(255) is a common choice for string types like names in PostgreSQL
    -- when you want to enforce a maximum length. TEXT is also valid for arbitrary length.
    first_name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    -- No NOT NULL for nickname, allowing it to be optional (can be NULL).
    nickname VARCHAR(255)
);