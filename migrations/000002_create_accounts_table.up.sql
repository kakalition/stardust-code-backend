CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  userId INTEGER,
  emoji VARCHAR,
  name VARCHAR,
  balance NUMERIC,
  createdDate TIMESTAMP,
  updatedDate TIMESTAMP
);