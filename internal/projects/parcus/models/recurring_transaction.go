package models

import (
	"stardustcode/backend/internal/utils"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type RecurringTransaction struct {
	Id                  *string          `db:"id"`
	UserId              *string          `db:"user_id"`
	LocalId             *string          `db:"local_id"`
	Name                *string          `db:"name"`
	Frequency           *string          `db:"frequency"`
	StartDate           pgtype.Date      `db:"start_date"`
	NextDueDate         pgtype.Date      `db:"next_due_date"`
	CategoryId          *int             `db:"category_id"`
	DeductFromAccountId *int             `db:"deduct_from_wallet_account_id"`
	Amount              *float64         `db:"amount"`
	Notes               *string          `db:"notes"`
	CreatedDate         pgtype.Timestamp `db:"created_date"`
	UpdatedDate         pgtype.Timestamp `db:"updated_date"`
}

type NetworkRecurringTransaction struct {
	Id                  *string    `json:"id"`
	LocalId             *string    `json:"localId" validate:"required"`
	Name                *string    `json:"name" validate:"required"`
	Frequency           *string    `json:"frequency" validate:"required"`
	StartDate           *time.Time `json:"startDate" validate:"required"`
	NextDueDate         *time.Time `json:"nextDueDate" validate:"required"`
	CategoryId          *int       `json:"categoryId" validate:"required"`
	DeductFromAccountId *int       `json:"deductFromAccountId" validate:"required"`
	Amount              *float64   `json:"amount" validate:"required"`
	Notes               *string    `json:"notes" validate:"required"`
	CreatedDate         *time.Time `json:"createdDate" validate:"required"`
	UpdatedDate         *time.Time `json:"updatedDate"`
}

func (u *RecurringTransaction) ToNetwork() NetworkRecurringTransaction {
	return NetworkRecurringTransaction{
		Id:                  u.Id,
		LocalId:             u.LocalId,
		Name:                u.Name,
		Frequency:           u.Frequency,
		StartDate:           &u.StartDate.Time,
		NextDueDate:         &u.NextDueDate.Time,
		CategoryId:          u.CategoryId,
		DeductFromAccountId: u.DeductFromAccountId,
		Amount:              u.Amount,
		Notes:               u.Notes,
		CreatedDate:         &u.CreatedDate.Time,
		UpdatedDate:         &u.UpdatedDate.Time,
	}
}

func (u *NetworkRecurringTransaction) ToInternal() RecurringTransaction {
	return RecurringTransaction{
		Id:        u.Id,
		LocalId:   u.LocalId,
		Name:      u.Name,
		Frequency: u.Frequency,
		StartDate: pgtype.Date{
			Time: *u.StartDate,
		},
		NextDueDate: pgtype.Date{
			Time: *u.NextDueDate,
		},
		CategoryId:          u.CategoryId,
		DeductFromAccountId: u.DeductFromAccountId,
		Amount:              u.Amount,
		Notes:               u.Notes,
		CreatedDate: pgtype.Timestamp{
			Time: *u.CreatedDate,
		},
		UpdatedDate: utils.PgTimestampGetter(u.UpdatedDate),
	}
}
