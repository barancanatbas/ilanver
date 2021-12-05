package request

type UserRegister struct {
	Name        string `validate:"required" json:"name"`
	Surname     string `validate:"required" json:"surname"`
	Phone       string `validate:"required" json:"phone"`
	Password    string `validate:"required" json:"password"`
	Email       string `validate:"required" json:"email"`
	Birthday    string `validate:"required" json:"birthday"`
	Districtfk  uint   `validate:"required" json:"districtfk"`
	Description string `validate:"required" json:"description"`
}

type UserLogin struct {
	Phone    string `validate:"required" json:"phone"`
	Password string `validate:"required" json:"password"`
}
