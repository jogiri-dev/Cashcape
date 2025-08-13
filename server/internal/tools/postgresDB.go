package tools

import (
	"log"
	"os"

	"github.com/jogiri-dev/cashcape/server/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

// GetExpenses implements DatabaseInterface.
func (p *postgresDB) GetExpenses(userId string) *api.ExpensesResponse {
	var expenses []api.Expense
	p.db.Find(&expenses) // TODO: Implement Multiuser lookup
	return &api.ExpensesResponse{Expenses: expenses}
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
