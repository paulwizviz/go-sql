CREATE TABLE IF NOT EXISTS person (
    -- SERIAL is the PostgreSQL-specific pseudo-type for auto-incrementing integer primary keys.
    -- It implicitly creates a sequence and sets a NOT NULL constraint.
    id SERIAL PRIMARY KEY
);

