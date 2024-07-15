CREATE TABLE IF NOT EXISTS transactions (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR,
  local_id VARCHAR,
  date TIMESTAMP,
  wallet_id INTEGER,
  category_id INTEGER,
  amount NUMERIC,
  notes VARCHAR,
  isRecurring BOOLEAN,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);