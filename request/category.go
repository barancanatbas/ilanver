package request

type CategoryInsert struct {
	CategoryName string `json:"category_name" validate:"required"`
	MainCategory uint   `json:"main_category_id"`
}

type CategoryUpdate struct {
	Id           uint   `json:"id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
	MainCategory uint   `json:"main_category_id" validate:"omitempty"`
}
