package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jogiri-dev/cashcape/server/api"
	"github.com/jogiri-dev/cashcape/server/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	var params = api.ExpensesParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Database calls and error handling here
	var database *tools.DatabaseInterface

	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// TODO: Retrieve this info from Auth token
	var expenseResponse = (*database).GetExpenses(1)

	if expenseResponse == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	var response = api.ExpensesResponse{
		Expenses: expenseResponse.Expenses,
		// CODE??
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
