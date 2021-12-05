package request

type CategoryInsert struct {
	CategoryName string `json:"category_name" validate:"required"`
	MainCategory uint   `json:"main_category_id"`
}
