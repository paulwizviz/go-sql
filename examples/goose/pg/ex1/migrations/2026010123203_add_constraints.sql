-- 20260101000003_add_constraints.sql
-- +goose Up
ALTER TABLE person_name_identifier
    ADD CONSTRAINT fk_person FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_named_id FOREIGN KEY (named_identifier_id) REFERENCES named_identifier(id) ON DELETE CASCADE,
    ADD CONSTRAINT uq_person_name UNIQUE (person_id, named_identifier_id);

-- +goose Down
ALTER TABLE person_name_identifier 
    DROP CONSTRAINT fk_person,
    DROP CONSTRAINT fk_named_id,
    DROP CONSTRAINT uq_person_name;