package request

type InsertProductState struct {
	State string `json:"state"`
}

type UpdateProductState struct {
	ID    uint   `json:"id"`
	State string `json:"state"`
}
