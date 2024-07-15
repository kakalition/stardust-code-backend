package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id           *string          `db:"id"`
	DisplayName  *string          `db:"display_name"`
	Email        *string          `db:"email"`
	LastSignedIn pgtype.Timestamp `db:"last_signed_in"`
}

type NetworkUser struct {
	Id           *string    `json:"id"`
	DisplayName  *string    `json:"displayName" validate:"required"`
	Email        *string    `json:"email" validate:"required"`
	LastSignedIn *time.Time `json:"lastSignedIn"`
}

func (u *User) ToNetwork() NetworkUser {
	return NetworkUser{
		Id:           u.Id,
		DisplayName:  u.DisplayName,
		Email:        u.Email,
		LastSignedIn: &u.LastSignedIn.Time,
	}
}

func (u *NetworkUser) ToInternal() User {
	return User{
		Id:          u.Id,
		DisplayName: u.DisplayName,
		Email:       u.Email,
	}
}
