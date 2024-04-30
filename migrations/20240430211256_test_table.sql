-- Up migration
-- Write SQL statements for applying the migration here.
CREATE TABLE IF NOT EXISTS test (
    id BYTEA PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Down migration
-- Write SQL statements for reverting the migration here.
DROP TABLE IF EXISTS test;