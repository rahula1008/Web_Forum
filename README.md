# Web Forum

Reddit-like web forum with a React (TypeScript) frontend and a Go (Gin) backend.

## Requirements
- Node.js 18+ and npm
- Go 1.21+
- PostgreSQL (local or hosted)

## Repository Structure
- `frontend/` - React + Vite app
- `backend/` - Go API server

## Backend Setup
1) Create `backend/.env`:

```bash
PORT=3000 # or port of your choice
SECRET=<Secret> # JWT secret
DB_URI=<SUPABASE_SESSION_POOLER_CONN_STRING>
# Ex: DB_URI=postgresql://<user>:[YOUR-PASSWORD]@host:<port>/<database>
```

2) Run database migrations:

```bash
cd backend

go run migrate/migrate.go
```

3) Start the API server:

```bash
cd backend

go run main.go
```

The server will start on the default Gin port (usually `:8080`).

## Frontend Setup
1) Install dependencies:

```bash
cd frontend

npm install
```

2) Create `frontend/.env`:

```bash
VITE_BACKEND_URL=<BACKEND_URL>
```

3) Start the dev server:

```bash
cd frontend

npm run dev
```

For more details on linting and building go to the README in the frontend folder

## Notes
- API routes are listed in `backend/README.md`.
- The frontend API client is configured in `frontend/src/auth/client.ts`.
- CORS is configured in `backend/main.go` for local and deployed frontend origins.
