CREATE TABLE IF NOT EXISTS person_name_identifier (
    id SERIAL PRIMARY KEY,
    -- Data type should match the referenced primary key in 'person' (INTEGER for SERIAL).
    person_id INTEGER NOT NULL,
    -- Data type should match the referenced primary key in 'named_identifier' (INTEGER for SERIAL).
    named_identifier_id INTEGER NOT NULL,

    -- Foreign key constraints are very similar to SQLite,
    -- using standard SQL syntax for REFERENCES and CASCADE actions.
    FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (named_identifier_id) REFERENCES named_identifier(id) ON DELETE CASCADE ON UPDATE CASCADE,

    -- The UNIQUE constraint on the pair of foreign keys ensures no duplicate associations.
    UNIQUE (person_id, named_identifier_id)
);