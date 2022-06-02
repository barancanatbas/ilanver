package request

type UpdateAddress struct {
	ID       uint   `json:"id"`
	District int16  `json:"district"`
	Detail   string `json:"detail"`
}
