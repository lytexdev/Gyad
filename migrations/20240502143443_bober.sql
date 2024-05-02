-- Up migration
-- Write SQL statements for applying the migration here.
CREATE TABLE IF NOT EXISTS bober (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Down migration
-- Write SQL statements for reverting the migration here.
DROP TABLE IF EXISTS bober;