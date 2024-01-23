ALTER TABLE messages
ADD CONSTRAINT unique_sender_receiver UNIQUE(sender_id, receiver_id, message_id);
