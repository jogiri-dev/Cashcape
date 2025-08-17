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

func (p *postgresDB) GetExpensesList(userId string, params api.GetExpensesParams) ([]api.Expense, error) {
	var expenses []api.Expense
	query := p.db.Preload("Category").Where(&api.Expense{UserID: userId})

	if !params.TimeStart.IsZero() {
		query = query.Where("date >= ?", params.TimeStart)
	}
	if !params.TimeEnd.IsZero() {
		query = query.Where("date <= ?", params.TimeEnd)
	}

	if err := query.Find(&expenses); err.Error != nil {
		log.Error(err.Error)
		return nil, err.Error
	}

	return expenses, nil
}

func (p *postgresDB) GetExpensesCategoryAggregates(userId string, params api.GetExpensesParams) ([]api.ExpenseCategoryAggregate, error) {
	var aggregates []api.ExpenseCategoryAggregate

	query := p.db.Table("expenses").
		Where("expenses.user_id = ?", userId).
		Joins("LEFT JOIN categories ON expenses.category_id = categories.id").
		Select("expenses.category_id,  categories.description as category_description, categories.symbol  as category_symbol, SUM(expenses.amount) as amount_sum").
		Group("expenses.category_id, categories.description, categories.symbol")

	if !params.TimeStart.IsZero() {
		query = query.Where("date >= ?", params.TimeStart)
	}
	if !params.TimeEnd.IsZero() {
		query = query.Where("date <= ?", params.TimeEnd)
	}

	if err := query.Scan(&aggregates); err.Error != nil {
		log.Error(err.Error)
		return nil, err.Error
	}

	return aggregates, nil
}

func (p *postgresDB) AddExpense(params api.AddExpenseParams) (*api.AddExpenseResponse, error) {

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
