CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  emoji VARCHAR,
  name VARCHAR,
  balance NUMERIC,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);