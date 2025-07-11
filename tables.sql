
CREATE TABLE collections (
    collection_id uuid DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    owner VARCHAR,
    PRIMARY KEY (collection_id)
);

CREATE INDEX collection ON collections (name);
CREATE INDEX owner ON collections (owner);

CREATE TABLE books (
    book_id uuid DEFAULT gen_random_uuid(),
    collection uuid,
    title VARCHAR NOT NULL,
    author VARCHAR NOT NULL,
    published VARCHAR NOT NULL,
    edition VARCHAR,
    description VARCHAR,
    genre VARCHAR,
    notes VARCHAR,
    PRIMARY KEY (book_id)
);

CREATE INDEX genres ON books (genre);
CREATE INDEX author ON books (author);