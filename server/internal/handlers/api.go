package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Use(chimiddle.BasicAuth("cashcape", map[string]string{"user": "123"}))

	r.Route("/expenses", func(router chi.Router) {
		router.Get("/", GetExpenses)
	})
}
