package tools

import (
	"github.com/jogiri-dev/cashcape/server/api"
	log "github.com/sirupsen/logrus"
)

type DatabaseInterface interface {
	GetExpenses(userId string) *api.ExpensesResponse
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
