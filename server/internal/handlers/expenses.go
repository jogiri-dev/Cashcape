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

	userId, ok := getUserIDfromContext(r)
	if !ok {
		log.Error("could not retrieve userId")
		api.InternalErrorHandler(w)
		return
	}

	if err := decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error(err)
		http.Error(w, "Invalid query parameters: "+err.Error(), http.StatusBadRequest)
		return
	}

	expenseResponse, err := h.db.GetExpensesList(userId, params)
	if err != nil {
		log.WithError(err).Error("Could not retrieve expense list from DB") // TODO: Check what happens if empty array?
		api.InternalErrorHandler(w)
		return
	}

	expenseAggregatesResponse, err := h.db.GetExpensesCategoryAggregates(userId, params)
	if err != nil {
		log.WithError(err).Error("Could not retrieve aggregated expenses per category from DB") // TODO: Check what happens if empty array?
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := api.GetExpensesResponse{
		Expenses:           expenseResponse,
		CategoryAggregates: expenseAggregatesResponse,
	}

	err = json.NewEncoder(w).Encode(response)

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

	if err := h.validator.Struct(req.Expense); err != nil {
		log.Error(err)
		api.ValidationErrorHandler(w, err)
		return
	}

	// Add userId from context
	userId, ok := getUserIDfromContext(r)
	if !ok {
		log.Error("Could not retrieve userId")
		api.InternalErrorHandler(w)
		return
	}
	req.Expense.UserID = userId

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

func getUserIDfromContext(r *http.Request) (string, bool) {
	userId, ok := r.Context().Value("userContextID").(string)

	return userId, ok
}
