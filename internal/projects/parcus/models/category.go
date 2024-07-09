package models

type Category struct {
	Id           *int    `json:"id" db:"id"`
	UserId       *string `json:"user_id" db:"user_id"`
	Emoji        *string `json:"emoji" db:"emoji"`
	Name         *string `json:"name" db:"name"`
	CategoryType *string `json:"category_type" db:"category_type"`
}
