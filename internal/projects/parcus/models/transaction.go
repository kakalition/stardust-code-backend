package models

import (
	"stardustcode/backend/internal/utils"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Transaction struct {
	Id          *string          `db:"id"`
	UserId      *string          `db:"user_id"`
	LocalId     *string          `db:"local_id"`
	Date        pgtype.Date      `db:"date"`
	AccountId   *int             `db:"account_id"`
	CategoryId  *int             `db:"category_id"`
	Amount      *float64         `db:"amount"`
	Notes       *string          `db:"notes"`
	IsRecurring *bool            `db:"is_recurring"`
	CreatedDate pgtype.Timestamp `db:"created_date"`
	UpdatedDate pgtype.Timestamp `db:"updated_date"`
}

type NetworkTransaction struct {
	Id          *string    `json:"id"`
	LocalId     *string    `json:"localId" validate:"required"`
	Date        *time.Time `json:"date" validate:"required"`
	AccountId   *int       `json:"accountId" validate:"required"`
	CategoryId  *int       `json:"categoryId" validate:"required"`
	Amount      *float64   `json:"amount" validate:"required"`
	Notes       *string    `json:"notes" validate:"required"`
	IsRecurring *bool      `json:"isRecurring" validate:"required"`
	CreatedDate *time.Time `json:"createdDate" validate:"required"`
	UpdatedDate *time.Time `json:"updatedDate"`
}

func (u *Transaction) ToNetwork() NetworkTransaction {
	return NetworkTransaction{
		Id:          u.Id,
		LocalId:     u.LocalId,
		Date:        &u.Date.Time,
		AccountId:   u.AccountId,
		CategoryId:  u.CategoryId,
		Amount:      u.Amount,
		Notes:       u.Notes,
		IsRecurring: u.IsRecurring,
		CreatedDate: &u.CreatedDate.Time,
		UpdatedDate: &u.UpdatedDate.Time,
	}
}

func (u *NetworkTransaction) ToInternal() Transaction {
	return Transaction{
		Id:      u.Id,
		LocalId: u.LocalId,
		Date: pgtype.Date{
			Time: *u.Date,
		},
		AccountId:   u.AccountId,
		CategoryId:  u.CategoryId,
		Amount:      u.Amount,
		Notes:       u.Notes,
		IsRecurring: u.IsRecurring,
		CreatedDate: pgtype.Timestamp{
			Time: *u.CreatedDate,
		},
		UpdatedDate: utils.PgTimestampGetter(u.UpdatedDate),
	}
}
