# Go Backend Migration (`backedn_go`)

This folder contains the migrated Node.js backend implemented in Go with:

- `gin` for HTTP/context
- `gorm` for PostgreSQL
- JWT auth middleware
- Single API port with service-based route groups (as requested)

## Architecture

Service segregation is done by folder + route group, while keeping one API process:

- `router/router.go`:
  - `auth` group (`/auth`)
  - `project` group (`/projects`)
  - `task` group (`/projects/:id/tasks`, `/tasks/:id`)

Layered folder design per service:

- `input_mapper/<service>`
- `controller/<service>`
- `usecase/<service>`
- `output_mapper/<service>`
- `repository/<service>`

Business logic lives in `usecase` only. Controllers only:
1. map/clean input
2. call usecase
3. map output response

## Endpoints (same surface)

- `POST /auth/register`
- `POST /auth/login`
- `GET /projects`
- `POST /projects`
- `GET /projects/:id`
- `PATCH /projects/:id`
- `DELETE /projects/:id`
- `GET /projects/:id/tasks`
- `POST /projects/:id/tasks`
- `PATCH /tasks/:id`
- `DELETE /tasks/:id`

## Run with Docker (no local Go SDK needed)

From `backedn_go`:

```bash
docker compose up --build
```

## Local env vars

Copy `.env.example` values into your environment:

- `DATABASE_URL`
- `JWT_SECRET`
- `PORT`

## Local setup without system Go install

This migration includes a local Go SDK at:

- `../_local_go_sdk`

To install dependencies and build:

```powershell
./setup.ps1
```

## GoLand one-click run

Open `backend_go` in GoLand and run [main.go](/C:/Users/naren/Downloads/Greening_India/backend_go/main.go) using the Run button.
It starts the API server and reads env from `.env`.
