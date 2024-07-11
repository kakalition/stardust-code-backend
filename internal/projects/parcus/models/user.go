package models

type User struct {
	Id          *int    `json:"id" db:"id"`
	DisplayName *string `json:"displayName" db:"display_name" validate:"required"`
	Email       *string `json:"email" db:"email" validate:"required"`
}
