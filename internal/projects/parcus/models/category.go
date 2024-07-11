package models

type Category struct {
	Id           *int    `json:"id" db:"id"`
	UserId       *string `json:"userId" db:"user_id"`
	Emoji        *string `json:"emoji" db:"emoji" validate:"required"`
	Name         *string `json:"name" db:"name" validate:"required"`
	CategoryType *string `json:"categoryType" db:"category_type" validate:"required"`
}
