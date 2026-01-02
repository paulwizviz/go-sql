-- 20260101000001_create_base_tables.sql
-- +goose Up
CREATE TABLE IF NOT EXISTS person (id SERIAL PRIMARY KEY);
CREATE TABLE IF NOT EXISTS named_identifier (id SERIAL PRIMARY KEY);
CREATE TABLE IF NOT EXISTS person_name_identifier (id SERIAL PRIMARY KEY);

-- +goose Down
DROP TABLE IF EXISTS person_name_identifier;
DROP TABLE IF EXISTS named_identifier;
DROP TABLE IF EXISTS person;
