CREATE TABLE IF NOT EXISTS budgets (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR,
  local_id VARCHAR,
  category_id INTEGER,
  amount NUMERIC,
  period DATE
);