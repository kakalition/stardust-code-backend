package models

import (
	"time"
)

type Transaction struct {
	Id          int
	Date        time.Time
	WalletId    *int
	CategoryId  *int
	Amount      float64
	Notes       string
	Type        string
	IsRecurring bool
	CreatedDate time.Time
	UpdatedDate *time.Time
}
