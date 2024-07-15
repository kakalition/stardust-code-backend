CREATE TABLE IF NOT EXISTS categories (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR,
  local_id VARCHAR,
  emoji VARCHAR,
  name VARCHAR,
  category_type VARCHAR
);