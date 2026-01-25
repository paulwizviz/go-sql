-- 20260101000001_create_base_tables.sql
-- +goose Up
CREATE TABLE IF NOT EXISTS person (id SERIAL PRIMARY KEY);
CREATE TABLE IF NOT EXISTS named_identifier (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    surname TEXT NOT NULL,
    nickname TEXT
);
CREATE TABLE IF NOT EXISTS person_name_identifier (
    id SERIAL PRIMARY KEY,
    person_id REFERENCES person(id) ON DELETE CASCADE,
    named_identifier_id REFERENCES named_identifier(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS person_name_identifier;
DROP TABLE IF EXISTS named_identifier;
DROP TABLE IF EXISTS person;
