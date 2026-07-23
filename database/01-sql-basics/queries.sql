-- Run these statements in PostgreSQL after creating a practice database.

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

-- CREATE: add rows.
INSERT INTO users (name, email)
VALUES
    ('Arpit', 'arpit@example.com'),
    ('Neha', 'neha@example.com');

-- READ: select the columns you need.
SELECT id, name, email
FROM users;

-- READ one row with a condition.
SELECT id, name, email
FROM users
WHERE email = 'arpit@example.com';

-- UPDATE: change only matching rows.
UPDATE users
SET name = 'Arpit Kuriyal'
WHERE id = 1;

-- DELETE: remove only matching rows.
DELETE FROM users
WHERE id = 2;
