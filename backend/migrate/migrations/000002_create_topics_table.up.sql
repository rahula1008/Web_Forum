CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,

    title CITEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    
    -- Foreign Key: Links to the Users table
    creator_id INT NOT NULL REFERENCES users(id) ON DELETE RESTRICT, 

    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE 
);
