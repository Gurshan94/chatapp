CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    room_id INT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    deleted BOOLEAN DEFAULT FALSE,  -- Soft delete support
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for optimization
CREATE INDEX idx_messages_room ON messages(room_id);
CREATE INDEX idx_messages_created ON messages(created_at);