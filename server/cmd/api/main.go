package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jogiri-dev/cashcape/server/internal/handlers"
	"github.com/jogiri-dev/cashcape/server/internal/tools"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
    	log.Println("No .env file found, relying on environment variables")
	}

	log.SetReportCaller(true)

	db, err := tools.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	h := handlers.NewHandler(db)
	h.RegisterRoutes(r)

	var port = ":8080"

	fmt.Printf("Starting GO API service on http://%v", port)

	if err = http.ListenAndServe(port, r); err != nil {
		log.Error(err)
	}
}
