CREATE TABLE IF NOT EXISTS recurring_transactions (
  id SERIAL PRIMARY KEY,
  userId INTEGER,
  name VARCHAR,
  frequency VARCHAR,
  startDate DATE,
  nextDueDate DATE,
  categoryId INTEGER,
  deductFromAccountId INTEGER,
  amount NUMERIC,
  notes VARCHAR,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);