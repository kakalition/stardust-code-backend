package models

import "time"

type Budget struct {
	Id         *string    `db:"id"`
	UserId     *string    `db:"user_id"`
	LocalId    *string    `db:"local_id"`
	CategoryId *string    `db:"categoryId"`
	Amount     *float64   `db:"amount"`
	Period     *time.Time `db:"period"`
}

type NetworkBudget struct {
	Id         *string    `json:"id"`
	LocalId    *string    `json:"localId"`
	CategoryId *string    `json:"categoryId" validate:"required"`
	Amount     *float64   `json:"amount" validate:"required"`
	Period     *time.Time `json:"period" validate:"required"`
}

func (u *Budget) ToNetwork() NetworkBudget {
	return NetworkBudget{
		Id:         u.Id,
		LocalId:    u.LocalId,
		CategoryId: u.CategoryId,
		Amount:     u.Amount,
		Period:     u.Period,
	}
}

func (u *NetworkBudget) ToInternal() Budget {
	return Budget{
		Id:         u.Id,
		LocalId:    u.LocalId,
		CategoryId: u.CategoryId,
		Amount:     u.Amount,
		Period:     u.Period,
	}
}
