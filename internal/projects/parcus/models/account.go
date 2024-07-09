package models

import "time"

type Account struct {
	Id          *int       `json:"id" db:"id"`
	UserId      *string    `json:"user_id" db:"user_id"`
	Emoji       *string    `json:"emoji" db:"emoji"`
	Name        *string    `json:"name" db:"name"`
	Balance     *float64   `json:"balance" db:"balance"`
	CreatedDate *time.Time `json:"created_date" db:"created_date"`
	UpdatedDate *time.Time `json:"updated_date" db:"updated_date"`
}
