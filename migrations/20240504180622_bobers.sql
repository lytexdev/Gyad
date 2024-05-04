-- Up migration
-- Write SQL statements for applying the migration here.
CREATE TABLE IF NOT EXISTS bobers (
    id UUID PRIMARY KEY UNIQUE,
    name VARCHAR(255),
    age INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO bobers (id, name, age) VALUES ('f7b3b3b4-3b7b-4b7b-8b7b-7b7b7b7b7b7b', 'bobersino', 12);
INSERT INTO bobers (id, name, age) VALUES ('f7b3b3b4-3b7b-4b7b-8b7b-7c7c7c7c7c7c', 'baber', 3);
INSERT INTO bobers (id, name, age) VALUES ('f7b3b3b4-3b7b-4b7b-8b7b-7d7d7d7d7d7d', 'boba', 16);

-- Down migration
-- Write SQL statements for reverting the migration here.
DROP TABLE IF EXISTS bobers;
