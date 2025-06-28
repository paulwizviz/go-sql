CREATE TABLE IF NOT EXISTS person (
    -- SERIAL is the PostgreSQL-specific pseudo-type for auto-incrementing integer primary keys.
    -- It implicitly creates a sequence and sets a NOT NULL constraint.
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS named_identifier (
    id SERIAL PRIMARY KEY,
    -- VARCHAR(255) is a common choice for string types like names in PostgreSQL
    -- when you want to enforce a maximum length. TEXT is also valid for arbitrary length.
    first_name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    -- No NOT NULL for nickname, allowing it to be optional (can be NULL).
    nickname VARCHAR(255)
);

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
