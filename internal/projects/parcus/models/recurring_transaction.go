package models

import (
	"time"
)

type RecurringTransaction struct {
	Id                  *int       `json:"id" db:"id"`
	UserId              *int       `json:"user_id" db:"user_id"`
	Name                *string    `json:"name" db:"name"`
	Frequency           *string    `json:"frequency" db:"frequency"`
	StartDate           *time.Time `json:"start_date" db:"start_date"`
	NextDueDate         *time.Time `json:"next_due_date" db:"next_due_date"`
	CategoryId          *int       `json:"category_id" db:"category_id"`
	DeductFromAccountId *int       `json:"deduct_from_account_id" db:"deduct_from_account_id"`
	Amount              *float64   `json:"amount" db:"amount"`
	Notes               *string    `json:"notes" db:"notes"`
	CreatedDate         *time.Time `json:"created_date" db:"created_date"`
	UpdatedDate         *time.Time `json:"updated_date" db:"updated_date"`
}
