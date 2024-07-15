CREATE TABLE IF NOT EXISTS recurring_transactions (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR,
  local_id VARCHAR,
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