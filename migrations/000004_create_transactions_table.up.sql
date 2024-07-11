CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  userId INTEGER,
  date TIMESTAMP,
  walletId INTEGER,
  categoryId INTEGER,
  amount NUMERIC,
  notes VARCHAR,
  isRecurring BOOLEAN,
  createdDate TIMESTAMP,
  updatedDate TIMESTAMP
);