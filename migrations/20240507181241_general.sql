-- Up migration
-- Write SQL statements for applying the migration here.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Down migration
-- Write SQL statements for reverting the migration here.
DROP EXTENSION IF EXISTS "uuid-ossp";