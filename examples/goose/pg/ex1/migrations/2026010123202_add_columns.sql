-- 20260101000002_add_identity_columns.sql
-- +goose Up
ALTER TABLE named_identifier 
    ADD COLUMN first_name VARCHAR(255) NOT NULL,
    ADD COLUMN surname VARCHAR(255) NOT NULL,
    ADD COLUMN nickname VARCHAR(255);

ALTER TABLE person_name_identifier
    ADD COLUMN person_id INTEGER NOT NULL,
    ADD COLUMN named_identifier_id INTEGER NOT NULL;

-- +goose Down
ALTER TABLE person_name_identifier DROP COLUMN person_id, DROP COLUMN named_identifier_id;
ALTER TABLE named_identifier DROP COLUMN first_name, DROP COLUMN surname, DROP COLUMN nickname;
