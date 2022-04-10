package request

type InsertCategory struct {
	Name         string `json:"name"`
	MainCategory uint   `json:"maincategory"`
}

type UpdateCategory struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	MainCategory uint   `json:"maincategory"`
}
