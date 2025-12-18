-- Users
ALTER TABLE users
ALTER COLUMN updated_at DROP DEFAULT;

-- Topics
ALTER TABLE topics
ALTER COLUMN updated_at DROP DEFAULT;

-- Posts
ALTER TABLE posts
ALTER COLUMN updated_at DROP DEFAULT;

-- Comments
ALTER TABLE comments
ALTER COLUMN updated_at DROP DEFAULT;
