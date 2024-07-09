CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  emoji VARCHAR,
  name VARCHAR,
  category_type VARCHAR
);