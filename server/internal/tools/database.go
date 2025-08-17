package tools

import (
	"github.com/jogiri-dev/cashcape/server/api"
	log "github.com/sirupsen/logrus"
)

type DatabaseInterface interface {
	GetExpensesList(userId string, params api.GetExpensesParams) ([]api.Expense, error)
	GetExpensesCategoryAggregates(userId string, params api.GetExpensesParams) ([]api.ExpenseCategoryAggregate, error)
	AddExpense(expense api.AddExpenseParams) (*api.AddExpenseResponse, error)
	SetupDatabase() error
}

func NewDatabase() (DatabaseInterface, error) {

	// var database DatabaseInterface = &mockDB{}
	var database DatabaseInterface = &postgresDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return database, nil
}
