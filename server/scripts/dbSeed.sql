-- Use a fixed UUID for the demo user
INSERT INTO public.users (id) VALUES
  ('11111111-1111-1111-1111-111111111111');

-- Seed categories for the user (reference the UUID)
INSERT INTO categories (id, user_id, description, symbol) VALUES
  (1, '11111111-1111-1111-1111-111111111111', 'Food', '🍔'),
  (2, '11111111-1111-1111-1111-111111111111', 'Transport', '🚗');

-- Seed expenses for the user (reference the UUID)
INSERT INTO expenses (id, user_id, amount, currency, description, category_id, created_at) VALUES
  (1, '11111111-1111-1111-1111-111111111111', 12.5, 'EUR', 'Lunch', 1, NOW()),
  (2, '11111111-1111-1111-1111-111111111111', 3.2, 'EUR', 'Bus ticket', 2, NOW());