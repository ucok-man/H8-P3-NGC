-- Insert records into the 'users' table
INSERT INTO users (username, email, password) VALUES
  ('JohnDoe', 'john.doe@example.com', 'securepass'),
  ('JaneSmith', 'jane.smith@example.com', 'password123'),
  ('AliceJohnson', 'alice.johnson@example.com', 'qwerty'),
  ('BobWilliams', 'bob.williams@example.com', 'pass456');

-- Insert records into the 'messages' table
INSERT INTO messages (sender_id, receiver_id, subject, body) VALUES
  (1, 2, 'Hello', 'Hi there! How are you?'),
  (2, 1, '', 'I am doing well, thanks! How about you?'),
  (1, 2, '', 'I am good too. Anything interesting happening?'),
  (2, 1, '', 'Not much, just planning for the weekend. You?'),
  (1, 2, '', 'Same here. Maybe we can catch up?'),
  (2, 1, 'Sure', 'Sounds like a plan! Where do you want to meet?');
