CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  category_id INTEGER,
  amount NUMERIC,
  period DATE
);