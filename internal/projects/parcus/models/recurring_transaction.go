package models

import (
	"time"
)

type RecurringTransaction struct {
	Id                  *int       `json:"id" db:"id"`
	UserId              *int       `json:"userId" db:"userId"`
	Name                *string    `json:"name" db:"name" validate:"required"`
	Frequency           *string    `json:"frequency" db:"frequency" validate:"required"`
	StartDate           *time.Time `json:"startDate" db:"startDate" validate:"required"`
	NextDueDate         *time.Time `json:"nextDueDate" db:"nextDueDate" validate:"required"`
	CategoryId          *int       `json:"categoryId" db:"categoryId" validate:"required"`
	DeductFromAccountId *int       `json:"deductFromAccountId" db:"deductFromAccountId" validate:"required"`
	Amount              *float64   `json:"amount" db:"amount" validate:"required"`
	Notes               *string    `json:"notes" db:"notes" validate:"required"`
	CreatedDate         *time.Time `json:"createdDate" db:"createdDate" validate:"required"`
	UpdatedDate         *time.Time `json:"updatedDate" db:"updatedDate"`
}
