# Web Forum Backend

Go + Gin backend for the Web Forum app. Provides REST endpoints for users, topics, posts, and comments, backed by PostgreSQL.

## Requirements
- Go 1.21+ (or your project Go version)
- PostgreSQL

## Setup
1) Configure environment variables:

This project uses supabase for postgres database hosting. For supabase connection, we are using the Session Pooler connection because it is IPv4 compatible. 

Create `backend/.env` with:

```bash
PORT=3000 # or port of your choice
SECRET=<Secret> # The secret used for hashing JWT 
DB_URI=<SUPABASE_SESSION_POOLER_CONN_STRING>
#Ex: DB_URI=postgresql://<user>:[YOUR-PASSWORD]@host:<port>/<database>
```

2) Run database migrations:

```bash
go run migrate/migrate.go
```

## Development
Run the API server:

```bash
go run main.go
```

The server will start on the default Gin port (usually `:8080`).

## API Endpoints


### Health
- `GET /ping` - health check

### Topics
- `GET /topics` - list topics
- `GET /topics/:id` - get topic by id
- `GET /topics/search?title=...` - search topics by title
- `POST /topics` - create topic
- `PUT /topics/:id` - update topic
- `DELETE /topics/:id` - delete topic

### Posts
- `GET /posts` - list posts
- `GET /posts/:id` - get post by id
- `GET /posts/search?title=...` - search posts by title
- `GET /topics/:id/posts` - list posts in a topic
- `POST /posts` - create post
- `PUT /posts/:id` - update post
- `DELETE /posts/:id` - delete post

### Comments
- `GET /comments` - list comments
- `GET /comments/:id` - get comment by id
- `GET /posts/:id/comments` - list comments for a post
- `POST /comments` - create comment
- `PUT /comments/:id` - update comment
- `DELETE /comments/:id` - delete comment

### Users / Auth
- `GET /users` - list users
- `GET /users/:id` - get user by id
- `GET /users/search?username=...` - search users by username
- `POST /users` - create user
- `POST /users/signup` - sign up
- `POST /users/login` - log in
- `POST /users/logout` - log out
- `GET /users/me` - current user (requires auth)
- `PUT /users/:id` - update user
- `DELETE /users/:id` - delete user

## Notes
- CORS is configured in `backend/main.go` for local and deployed frontend origins.
- Routes are defined under `backend/routes/` and handlers live in `backend/controllers/`.
