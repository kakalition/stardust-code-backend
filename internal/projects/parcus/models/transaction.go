package models

import (
	"time"
)

type Transaction struct {
	Id          *int       `json:"id" db:"id"`
	UserId      *int       `json:"userId" db:"userId"`
	Date        *time.Time `json:"date" db:"date" validate:"required"`
	WalletId    *int       `json:"walletId" db:"walletId" validate:"required"`
	CategoryId  *int       `json:"categoryId" db:"categoryId" validate:"required"`
	Amount      *float64   `json:"amount" db:"amount" validate:"required"`
	Notes       *string    `json:"notes" db:"notes" validate:"required"`
	IsRecurring *bool      `json:"isRecurring" db:"isRecurring" validate:"required"`
	CreatedDate *time.Time `json:"createdDate" db:"createdDate" validate:"required"`
	UpdatedDate *time.Time `json:"updatedDate" db:"updatedDate"`
}
