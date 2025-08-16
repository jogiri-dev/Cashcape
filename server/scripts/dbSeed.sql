BEGIN;

-- Use a fixed UUID for the demo user
INSERT INTO users (id, email) VALUES
  ('11111111-1111-1111-1111-111111111111', 'test@test.se');

-- Seed categories for the user (reference the UUID)
INSERT INTO categories (id, description, symbol) VALUES
  (1, 'Food', '🍔'),
  (2, 'Transport', '🚗');

-- Seed expenses for the user (reference the UUID)
INSERT INTO expenses (user_id, amount, currency, description, date, category_id) VALUES
  ('11111111-1111-1111-1111-111111111111', 125, 'SEK', 'Lunch', NOW(), 1),
  ('11111111-1111-1111-1111-111111111111', 32, 'SEK', 'Bus ticket', NOW(), 2);

  COMMIT;