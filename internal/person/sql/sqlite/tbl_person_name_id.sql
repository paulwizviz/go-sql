CREATE TABLE IF NOT EXISTS person_name_identifier (
    id INTEGER PRIMARY KEY, -- Primary key for the association record itself
    person_id INTEGER NOT NULL, -- Foreign key to the 'person' table, cannot be NULL
    named_identifier_id INTEGER NOT NULL, -- Foreign key to the 'named_identifier' table, cannot be NULL
    FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (named_identifier_id) REFERENCES named_identifier(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (person_id, named_identifier_id) -- Ensures that a person can only be linked to a specific named_identifier once
);
