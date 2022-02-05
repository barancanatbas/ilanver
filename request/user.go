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

type UserUpdate struct {
	Name        string `validate:"omitempty" json:"name"`
	Surname     string `validate:"omitempty" json:"surname"`
	Phone       string `validate:"omitempty" json:"phone"`
	Email       string `validate:"omitempty" json:"email"`
	Birthday    string `validate:"omitempty" json:"birthday"`
	Districtfk  uint   `validate:"omitempty" json:"districtfk"`
	Description string `validate:"omitempty" json:"description"`
}
