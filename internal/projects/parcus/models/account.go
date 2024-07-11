package models

import "time"

type Account struct {
	Id          *int       `json:"id" db:"id"`
	UserId      *string    `json:"userId" db:"user_id"`
	Emoji       *string    `json:"emoji" db:"emoji" validate:"required"`
	Name        *string    `json:"name" db:"name" validate:"required"`
	Balance     *float64   `json:"balance" db:"balance" validate:"required"`
	CreatedDate *time.Time `json:"createdDate" db:"createdDate" validate:"required"`
	UpdatedDate *time.Time `json:"updatedDate" db:"updatedDate"`
}
