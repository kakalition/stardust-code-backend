package models

import (
	"time"
)

type Transaction struct {
	Id          *int       `json:"id" db:"id"`
	UserId      *int       `json:"user_id" db:"user_id"`
	Date        *time.Time `json:"date" db:"date"`
	WalletId    *int       `json:"wallet_id" db:"wallet_id"`
	CategoryId  *int       `json:"category_id" db:"category_id"`
	Amount      *float64   `json:"amount" db:"amount"`
	Notes       *string    `json:"notes" db:"notes"`
	IsRecurring *bool      `json:"is_recurring" db:"is_recurring"`
	CreatedDate *time.Time `json:"created_date" db:"created_date"`
	UpdatedDate *time.Time `json:"updated_date" db:"updated_date"`
}
