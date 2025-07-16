CREATE TABLE IF NOT EXISTS named_identifier (
    id INTEGER PRIMARY KEY, -- Added PRIMARY KEY to make 'id' the primary key and auto-incrementing
    first_name TEXT NOT NULL, -- Added NOT NULL as names are typically required
    surname TEXT NOT NULL -- Added NOT NULL as names are typically required
    nickname TEXT -- Some case there could be no nickname
);
