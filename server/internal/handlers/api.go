package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/jogiri-dev/cashcape/server/api"
	"github.com/jogiri-dev/cashcape/server/internal/tools"
)

type Handler struct {
	db        tools.DatabaseInterface
	validator *validator.Validate
}

func NewHandler(db tools.DatabaseInterface) *Handler {
	v := validator.New()
	v.RegisterValidation("currency", api.CurrencyValidationFn)
	return &Handler{db: db, validator: v}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	// TODO: Implement auth
	// var apiUser, okUser = os.LookupEnv("APIuser")
	// var apiPassword, okPassword = os.LookupEnv("APIpassword")
	// if !okUser || !okPassword {
	// 	log.Fatal("Set API credentials in .env file")
	// }
	r.Use(chimiddle.WithValue("userContextID", "11111111-1111-1111-1111-111111111111")) // Hardcoded userId for MVP
	// r.Use(chimiddle.BasicAuth("cashcape", map[string]string{apiUser: apiPassword}))

	r.Route("/api/expenses", func(router chi.Router) {
		router.Get("/", h.GetExpenses)
		router.Post("/", h.AddExpense)
		router.Delete("/{id}", h.DeleteExpense)
	})

}
