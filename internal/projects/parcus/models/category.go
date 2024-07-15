package models

type Category struct {
	Id           *string `db:"id"`
	UserId       *string `db:"user_id"`
	LocalId      *string `db:"local_id"`
	Emoji        *string `db:"emoji"`
	Name         *string `db:"name"`
	CategoryType *string `db:"category_type"`
}

type NetworkCategory struct {
	Id           *string `json:"id"`
	LocalId      *string `json:"localId" validate:"required"`
	Emoji        *string `json:"emoji" validate:"required"`
	Name         *string `json:"name" validate:"required"`
	CategoryType *string `json:"categoryType" validate:"required"`
}

func (u *Category) ToNetwork() NetworkCategory {
	return NetworkCategory{
		Id:           u.Id,
		LocalId:      u.LocalId,
		Emoji:        u.Emoji,
		Name:         u.Name,
		CategoryType: u.CategoryType,
	}
}

func (u *NetworkCategory) ToInternal() Category {
	return Category{
		Id:           u.Id,
		LocalId:      u.LocalId,
		Emoji:        u.Emoji,
		Name:         u.Name,
		CategoryType: u.CategoryType,
	}
}
