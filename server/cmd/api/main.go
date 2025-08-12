package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jogiri-dev/cashcape/server/internal/handlers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	var port = "localhost:8000"

	fmt.Printf("Starting GO API service on http://%v", port)

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Error(err)
	}
}
