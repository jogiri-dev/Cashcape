package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Users struct {
	ID        int64     `json:"id"`        // Unique identifier
	Email     string    `json:"email"`     // Email serves as username
	CreatedAt time.Time `json:"createdAt"` // Account creation time
}

type Expense struct {
	ID          int64     `json:"id"` // Unique identifier
	UserID      int64     `json:"-"`  // Don't return
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Category struct {
	ID          int64  `json:"id"` // Unique identifier
	UserID      int64  `json:"-"`  // For future multi-user support
	Description string `json:"description"`
	Symbol      string `json:"symbol"` // Unicode symbol
}

type APIError struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type ExpensesResponse struct {
	Expenses []Expense `json:"expenses"` // Slice of all expenses
}

type ExpensesParams struct {
	// UserEmail string
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler  = func(w http.ResponseWriter, err error) { writeError(w, err.Error(), http.StatusBadRequest) }
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error ocurred", http.StatusInternalServerError)
	}
)
