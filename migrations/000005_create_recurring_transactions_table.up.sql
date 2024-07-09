CREATE TABLE IF NOT EXISTS recurring_transactions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  name VARCHAR,
  frequency VARCHAR,
  start_date DATE,
  next_due_date DATE,
  category_id INTEGER,
  deduct_from_account_id INTEGER,
  amount NUMERIC,
  notes VARCHAR,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);