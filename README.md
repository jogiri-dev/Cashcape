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

## Getting Started

1. **Create local Postgres DB with Docker Compose**

   ```sh
   docker compose up -d
   ```

2. **Create a `.env` file for environment variables**

   Create a `.env` file in the project root with the following content (replace with your own credentials as needed):

   ```
   DB_DSN="postgres://admin:examplepassword@localhost:5432/cashcape?sslmode=disable"
   ```

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
- [ ] Integration tests and CI/CD pipeline
- [ ] Dockerize backend and frontend for easy deployment
- [ ] Advanced dashboard analytics (category breakdowns, trends)
- [ ] Editable expenses and category management
- [ ] Mobile-friendly UI
- [ ] Offline mode
- [ ] API documentation and OpenAPI spec

---
