package handlers

import (
	"os"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/jogiri-dev/cashcape/server/internal/tools"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	DB tools.DatabaseInterface
}

func NewHandler(db tools.DatabaseInterface) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	var apiUser, okUser = os.LookupEnv("APIuser")
	var apiPassword, okPassword = os.LookupEnv("APIpassword")
	if !okUser || !okPassword {
		log.Fatal("Set API credentials in .env file")
	}

	r.Use(chimiddle.BasicAuth("cashcape", map[string]string{apiUser: apiPassword}))

	r.Route("/api/expenses", func(router chi.Router) {
		router.Get("/", h.GetExpenses)
	})
}
