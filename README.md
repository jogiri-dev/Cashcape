Create local Postgres db with docker compose
`docker compose up -d`

Run server to create tables with automigrate
`go run cmd/api/main.go`

Seed database using postgres CLI:
`psql -h localhost -U admin -d cashcape -f ./server/scripts/dbSeed.sql`
