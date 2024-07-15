package models

import (
	"stardustcode/backend/internal/utils"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	Id          *string          `db:"id"`
	UserId      *string          `db:"user_id"`
	LocalId     *string          `db:"local_id"`
	Emoji       *string          `db:"emoji"`
	Name        *string          `db:"name"`
	Balance     *float64         `db:"balance"`
	CreatedDate pgtype.Timestamp `db:"created_date"`
	UpdatedDate pgtype.Timestamp `db:"updated_date"`
}

type NetworkAccount struct {
	Id          *string    `json:"id"`
	LocalId     *string    `json:"localId" validate:"required"`
	Emoji       *string    `json:"emoji" validate:"required"`
	Name        *string    `json:"name" validate:"required"`
	Balance     *float64   `json:"balance" validate:"required"`
	CreatedDate *time.Time `json:"createdDate" validate:"required"`
	UpdatedDate *time.Time `json:"updatedDate"`
}

func (u *Account) ToNetwork() NetworkAccount {
	return NetworkAccount{
		Id:          u.Id,
		LocalId:     u.LocalId,
		Emoji:       u.Emoji,
		Name:        u.Name,
		Balance:     u.Balance,
		CreatedDate: &u.CreatedDate.Time,
		UpdatedDate: &u.UpdatedDate.Time,
	}
}

func (u *NetworkAccount) ToInternal() Account {
	return Account{
		Id:      u.Id,
		LocalId: u.LocalId,
		Emoji:   u.Emoji,
		Name:    u.Name,
		Balance: u.Balance,
		CreatedDate: pgtype.Timestamp{
			Time: *u.CreatedDate,
		},
		UpdatedDate: utils.PgTimestampGetter(u.UpdatedDate),
	}
}
