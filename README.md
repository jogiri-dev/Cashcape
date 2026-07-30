# Cashcape Expense Tracker

Cashcape is an expense tracker app built to help users manage and visualize their personal spending. The project is designed for easy extension and future multi-user support, with a clean separation between backend and frontend.

## Tech Stack

- **Backend:** Go
  - Organized with structs and interfaces
  - Functional principles for isolated logic
  - Uses [GORM](https://gorm.io/) for ORM and database migrations
  - Uses [Chi](https://github.com/go-chi/chi) for lightweight HTTP routing
  - Additional libraries: [logrus](https://github.com/sirupsen/logrus) for logging, [gorilla/schema](https://github.com/gorilla/schema) for request parsing, [go-playground/validator](https://github.com/go-playground/validator) for validation
- **Database:** PostgreSQL
- **Frontend:** React app created with Vite based on Material UI (MUI) Dashboard template

## Project Structure

- `server/` — Go backend
- `client/` — React frontend (Vite + MUI)
- `docker-compose.dev.yml` — local Postgres only, for local development
- `docker-compose.staging.yml` / `docker-compose.prod.yml` — full stack (db + backend + frontend) for Pi deployment
- `.github/workflows/` — CI/CD pipeline (build, push, staging auto-deploy, manual prod promotion)

## Getting Started

1. **Create local Postgres DB with Docker Compose**

   ```sh
   docker compose -f docker-compose.dev.yml up -d
   ```

2. **Create a `.env` file for environment variables**

   Copy the example file and fill in your own values:

   ```sh
   cp .env.example .env
   ```

   `.env.example` documents the required variables:

3. **Run the Go server to create tables with automigrate**

   ```sh
   go run cmd/api/main.go
   ```

4. **Seed the database using Postgres CLI**

   ```sh
   psql -h localhost -U admin -d cashcape -f ./server/scripts/dbSeed.sql
   ```

5. **Run frontend**
   ```sh
   npm run dev
   ```

## Deployment

The app runs on a self-hosted Raspberry Pi via Docker Compose, with images built and published automatically by GitHub Actions.

- **Registry:** Images are pushed to [GHCR](https://ghcr.io) as `expense-backend` and `expense-frontend`.
- **Staging:** Every push to `main` builds and tags images as `:staging`, which auto-deploys.
- **Production:** Promotion to `:prod` is a manual, approval-gated step in GitHub Actions — no rebuild, just retagging the exact image that was tested in staging.

To deploy manually on the Pi:

```sh
# staging
docker compose --env-file .env.staging -f docker-compose.staging.yml pull
docker compose --env-file .env.staging -f docker-compose.staging.yml up -d

# production
docker compose --env-file .env.prod -f docker-compose.prod.yml pull
docker compose --env-file .env.prod -f docker-compose.prod.yml up -d
```

`.env.prod` and `.env.staging` are created directly on the Pi (not committed to the repo) — same variables as `.env.example`, but with `db` as the Postgres host instead of `localhost`, and a distinct `POSTGRES_DB` per environment (e.g. `cashcape_staging`) to keep data isolated.

See `.github/workflows/build-and-deploy.yml` for the full pipeline.

## MVP Goals

- [x] **Get familiar with Go (my first GO project)**
- [x] **Single user support** (code prepared for multi-user expansion)
- [x] **Two main pages:**
  - **Expense list:** User can see expenses in list and add new expenses
  - **Dashboard & List:** User can view a dashboard (with charts) and a list of expenses, and delete expenses
- [x] **Postgres-backed persistent storage**
- [x] **Clean, modern UI with MUI dashboard**
- [x] **Go backend with clear error handling and code organization**

## Further Development

- [ ] Filtering by date and automatic grouping by month
- [ ] Multi-user authentication and authorization
- [ ] Integration tests
- [x] CI/CD pipeline (build, push, staging auto-deploy, manual prod promotion)
- [x] Dockerize backend and frontend for easy deployment
- [ ] Advanced dashboard analytics (category breakdowns, trends)
- [ ] Editable expenses and category management
- [ ] Mobile-friendly UI
- [ ] Offline mode
- [ ] API documentation and OpenAPI spec

---
