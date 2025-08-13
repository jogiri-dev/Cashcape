package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jogiri-dev/cashcape/server/api"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	var params = api.ExpensesParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	err := decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// TODO: Retrieve user ID from auth token/session
	expenseResponse := h.DB.GetExpenses("") // TODO: MVP getexpenses lists all expenses

	if expenseResponse == nil {
		log.Error(err) // TODO: Check what happens if empty array?
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
