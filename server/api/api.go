package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` // UUID as string
	Email     string    `json:"email" gorm:"not null"`                                                                                   // Email serves as username
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`                                                                         // Account creation time
}

type Expense struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	UserID      string    `json:"-" gorm:"type:uuid;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` // UUID as string, don't return in API
	User        User      `json:"-" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Amount      int64     `json:"amount" gorm:"not null" validate:"required,gt=0"` // In cents
	Currency    string    `json:"currency" gorm:"not null" validate:"currency"`    // Custom validation function to check for supported currencies                        // Prepare for later implementation, not required for now
	Description string    `json:"description" gorm:"not null" validate:"required"`
	CategoryID  *int64    `json:"categoryId" gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE;"` // Prepare for later implementation, empty for uncategorized
	Category    *Category `json:"category,omitempty" gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE;"`
	Date        time.Time `json:"date" gorm:"not null" validate:"required"` // Expense date. Set to midnight.
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type Category struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	UserID      *string   `json:"-" gorm:"type:uuid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` // UUID as string, don't return in API
	User        *User     `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Description string    `json:"description" gorm:"not null"`
	Symbol      string    `json:"symbol"` // Unicode symbol
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type ExpenseCategoryAggregate struct {
	CategoryID          *int64  `json:"categoryId"`
	CategoryDescription *string `json:"categoryDescription"`
	CategorySymbol      *string `json:"categorySymbol"`
	AmountSum           int64   `json:"amountSum"`
}

type GetExpensesResponse struct {
	Expenses           []Expense                  `json:"expenses"`
	CategoryAggregates []ExpenseCategoryAggregate `json:"categoryAggregates"`
}

type GetExpensesParams struct {
	TimeStart time.Time `schema:"timeStart"`
	TimeEnd   time.Time `schema:"timeEnd"`
}

type AddExpenseParams struct {
	Expense Expense
}
type AddExpenseResponse struct {
	Expense Expense `json:"expense"`
}

type ErrorResponse struct {
	Error   string       `json:"error"`
	Details []FieldError `json:"details,omitempty"`
}

type FieldError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var supportedCurrencies = map[string]struct{}{"SEK": {}}

func CurrencyValidationFn(fl validator.FieldLevel) bool {

	if _, ok := supportedCurrencies[fl.Field().String()]; !ok {
		return false
	}

	return true
}

func returnError(w http.ResponseWriter, message string, code int, details []FieldError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := ErrorResponse{
		Error:   message,
		Details: details,
	}

	json.NewEncoder(w).Encode(resp)
}

var (
	DecodeErrorHandler = func(w http.ResponseWriter, err error) {
		returnError(w, "Decoding error", http.StatusBadRequest, []FieldError{{Message: err.Error()}})
	}
	ValidationErrorHandler = func(w http.ResponseWriter, err error) {
		returnError(w, "Input validation failed", http.StatusBadRequest, parseValidationErrors(err))
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		returnError(w, "An unexpected error ocurred", http.StatusInternalServerError, []FieldError{})
	}
)

func parseValidationErrors(err error) []FieldError {
	var errs []FieldError
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, FieldError{
			Field:   e.Field(),
			Message: fmt.Sprintf("failed on '%s' rule", e.Tag()),
		})
	}
	return errs
}
