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

	fmt.Println("Starting GO API service...")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
