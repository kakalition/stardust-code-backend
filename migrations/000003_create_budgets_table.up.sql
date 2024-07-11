CREATE TABLE IF NOT EXISTS budgets (
  id SERIAL PRIMARY KEY,
  userId INTEGER,
  categoryId INTEGER,
  amount NUMERIC,
  period DATE
);