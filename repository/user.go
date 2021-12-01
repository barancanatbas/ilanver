package repository

import (
	"ilanver/internal/models"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (rootRepo *Repositories) User() UserRepo {
	return UserRepo{db: rootRepo.Db}
}

// login..
func (u UserRepo) Login(user models.User) (models.User, error) {
	err := u.db.Model(&models.User{}).Where("id = ? and phone = ?", user.ID, user.Phone).Take(&user).Error
	return user, err
}

// register new user..
func (u UserRepo) Register(user models.User) error {
	err := u.db.Model(&models.User{}).Save(&user)
	return err.Error
}

func (u UserRepo) ExistsPhone(phone string) bool {
	err := u.db.Model(&models.User{}).Where("phone = ?", phone)
	if err.Error != nil {
		return true
	} else if err.RowsAffected > 0 {
		return true
	}
	return false
}
