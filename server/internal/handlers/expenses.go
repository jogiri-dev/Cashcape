package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jogiri-dev/cashcape/server/api"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	var params = api.GetExpensesParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	if err := decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error(err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Retrieve user ID from auth token/session
	// TODO: MVP getexpenses lists all expenses
	expenseResponse, err := h.db.GetExpenses("")
	if err != nil {
		log.WithError(err).Error("Could not retrieve expenses from DB") // TODO: Check what happens if empty array?
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(expenseResponse)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

func (h *Handler) AddExpense(w http.ResponseWriter, r *http.Request) {
	var req = api.AddExpenseParams{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error(err)
		api.DecodeErrorHandler(w, err)
		return
	}

	log.Info(req)

	if err := h.validator.Struct(req.Expense); err != nil {
		log.Error(err)
		api.ValidationErrorHandler(w, err)
		return
	}

	response, err := h.db.AddExpense(req)

	if err != nil {
		log.WithError(err).Error("Could not add expense to DB")
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
