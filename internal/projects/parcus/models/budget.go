package models

import "time"

type Budget struct {
	Id         *int       `json:"id" db:"id"`
	UserId     *int       `json:"user_id" db:"user_id"`
	CategoryId *string    `json:"category_id" db:"category_id"`
	Amount     *float64   `json:"amount" db:"amount"`
	Period     *time.Time `json:"period" db:"period"`
}
