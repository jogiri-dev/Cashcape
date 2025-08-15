package tools

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jogiri-dev/cashcape/server/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

func (p *postgresDB) GetExpenses(userId string) (*api.GetExpensesResponse, error) {
	var expenses []api.Expense
	// TODO: Implement Multiuser lookup
	if result := p.db.Find(&expenses); result.Error != nil {
		return nil, result.Error
	}

	return &api.GetExpensesResponse{Expenses: expenses}, nil
}

func (p *postgresDB) AddExpense(params api.AddExpenseParams) (*api.AddExpenseResponse, error) {

	//TODO: Replace with real logic
	params.Expense.UserID = "11111111-1111-1111-1111-111111111111"

	if result := p.db.Create(&params.Expense); result.Error != nil {
		return nil, result.Error
	}

	return &api.AddExpenseResponse{Expense: params.Expense}, nil
}

func (p *postgresDB) SetupDatabase() error {
	var dsn, okDSN = os.LookupEnv("DB_DSN")
	if !okDSN {
		log.Fatal("Provide DB connection string in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Enable uuid-ossp extension for UUID generation
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&api.User{}, &api.Expense{}, &api.Category{}); err != nil {
		return err
	}

	p.db = db

	return nil
}
