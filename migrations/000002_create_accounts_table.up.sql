CREATE TABLE IF NOT EXISTS accounts (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR,
  local_id VARCHAR,
  emoji VARCHAR,
  name VARCHAR,
  balance NUMERIC,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);