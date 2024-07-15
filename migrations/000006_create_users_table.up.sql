CREATE TABLE IF NOT EXISTS users (
  id VARCHAR PRIMARY KEY,
  email VARCHAR,
  display_name VARCHAR,
  last_signed_in TIMESTAMP
);