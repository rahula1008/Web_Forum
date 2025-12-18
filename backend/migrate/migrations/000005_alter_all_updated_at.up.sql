--Users
ALTER TABLE users
ALTER COLUMN updated_at SET DEFAULT now();

--Topics
ALTER TABLE topics
ALTER COLUMN updated_at SET DEFAULT now();

--Posts
ALTER TABLE posts
ALTER COLUMN updated_at SET DEFAULT now();

--Comments 
ALTER TABLE comments
ALTER COLUMN updated_at SET DEFAULT now();
