-- Up migration
-- Write SQL statements for applying the migration here.
CREATE TABLE IF NOT EXISTS bober (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    age INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Down migration
-- Write SQL statements for reverting the migration here.
DROP TABLE IF EXISTS bober;
