CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  date TIMESTAMP,
  wallet_id INTEGER,
  category_id INTEGER,
  amount NUMERIC,
  notes VARCHAR,
  is_recurring BOOLEAN,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);