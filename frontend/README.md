# Web Forum Frontend

React + TypeScript + Vite frontend for the Web Forum app. It renders topics, posts, and comments, and integrates with the backend API for auth and CRUD actions.

## Requirements
- Node.js 18+ (or the version you use for the repo)
- npm

## Setup
1) Install dependencies:

```bash
npm install
```

2) Configure the backend URL:

Create `frontend/.env` with:

```bash
VITE_BACKEND_URL=<Insert URL Here>
# For example can use localhost:3000 or the deployment url of the API
```

## Development
Run the dev server:

```bash
npm run dev
```

## Build

```bash
npm run build
```

## Lint

```bash
npm run lint
```

## Notes
- API requests are configured in `frontend/src/auth/client.ts` and use `VITE_BACKEND_URL`.
- Auth relies on cookies (`withCredentials: true`), so your backend should set CORS accordingly.
