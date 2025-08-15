BEGIN;

-- Use a fixed UUID for the demo user
INSERT INTO public.users (id, email) VALUES
  ('11111111-1111-1111-1111-111111111111', 'test@test.se');

-- Seed categories for the user (reference the UUID)
INSERT INTO categories (description, symbol) VALUES
  ( 'Food', '🍔'),
  ( 'Transport', '🚗');

-- Seed expenses for the user (reference the UUID)
INSERT INTO expenses (user_id, amount, currency, description, date) VALUES
  ('11111111-1111-1111-1111-111111111111', 125, 'SEK', 'Lunch', NOW()),
  ('11111111-1111-1111-1111-111111111111', 32, 'SEK', 'Bus ticket', NOW());

  COMMIT;