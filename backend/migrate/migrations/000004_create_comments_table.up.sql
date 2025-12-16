CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,

    body TEXT NOT NULL,
    
    -- Foreign Key: Links to the Users table
	post_id INT NOT NULL REFERENCES posts(id) ON DELETE RESTRICT,
    creator_id INT NOT NULL REFERENCES users(id) ON DELETE RESTRICT, 

    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE 
);
