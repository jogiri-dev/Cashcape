package tools

import (
	"fmt"
	"time"

	"github.com/jogiri-dev/cashcape/server/api"
)

type mockDB struct{}

var mockExpenses = []api.Expense{
	{ID: 1,
		UserID:      1,
		Amount:      12.34,
		Currency:    "USD",
		Description: "Dummy expense User 1",
		CategoryID:  1,
		CreatedAt:   time.Now()},

	{ID: 2,
		UserID:      1,
		Amount:      12.34,
		Currency:    "USD",
		Description: "Dummy expense User 2",
		CategoryID:  1,
		CreatedAt:   time.Now()},
}

func (d *mockDB) GetExpenses(userId int64) *api.ExpensesResponse {

	// TODO: Filter out expenses here
	var clientData = api.ExpensesResponse{Expenses: mockExpenses}
	// clientData, ok := mockExpenses[username]

	// if !ok {
	// 	return nil
	// }

	return &clientData
}

// TODO: Remove auth logic?
var mockLoginDetails = map[string]LoginDetails{
	"testuser": {AuthToken: "123", Username: "testuser"},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	fmt.Println(clientData)

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
