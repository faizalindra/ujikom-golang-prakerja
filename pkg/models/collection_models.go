package models

type Collection struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Public      bool   `json:"public"`
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
}

type CollectionCreate struct {
	Name        string `json:"name" form:"name" valid:"required,string,length(1|255)"`
	Public      bool   `json:"public" form:"public" valid:"required,bool"`
	Description string `json:"description" form:"description" valid:"required,string,length(1|255)"`
}

type CollectionUpdate struct {
	Name        string `json:"name" form:"name" valid:"string,length(1|255)"`
	Public      bool   `json:"public" form:"public" valid:"bool"`
	Description string `json:"description" form:"description" valid:"string,length(1|255)"`
}

type CollectionResource struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Public      bool             `json:"public"`
	UserID      uint             `json:"user_id"`
	Description string           `json:"description"`
	RecipeCount uint             `json:"recipe_count"`
	Recipes     []RecipeResource `json:"recipes"`
}

type CollectionRecipesPivot struct {
	CollectionID uint `json:"collection_id" gorm:"foreignkey:ID;association_foreignkey:ID"`
	RecipeID     uint `json:"recipe_id" gorm:"foreignkey:ID;association_foreignkey:ID"`
}
