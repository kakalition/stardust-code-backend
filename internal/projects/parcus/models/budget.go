package models

import "time"

type Budget struct {
	Id         *int       `json:"id" db:"id"`
	CategoryId *string    `json:"categoryId" db:"categoryId" validate:"required"`
	Amount     *float64   `json:"amount" db:"amount" validate:"required"`
	Period     *time.Time `json:"period" db:"period" validate:"required"`
}
