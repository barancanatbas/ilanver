package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IUserRepo interface {
	Login(phone string) (model.User, error)
	Save(user *model.User) error
	Get(id uint) (model.User, error)
}

type UserRepo struct {
	tx *gorm.DB
}

// Compile time checks to ensure your type satisfies an interface
var _ IUserRepo = UserRepo{}

func NewUserRepository(tx *gorm.DB) UserRepo {
	return UserRepo{
		tx: tx,
	}
}

func (u UserRepo) Login(phone string) (model.User, error) {
	var user model.User
	err := u.tx.Model(&model.User{}).Where("phone = ?", phone).First(&user).Error

	return user, err

}

func (u UserRepo) Save(user *model.User) error {
	err := u.tx.Save(user).Error

	return err
}

func (u UserRepo) Get(id uint) (model.User, error) {
	var user model.User

	err := u.tx.Model(&model.User{}).Where("id = ?", id).First(&user).Error

	return user, err
}
