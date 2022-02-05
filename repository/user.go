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
func (u UserRepo) Login(user *models.User) error {
	err := u.db.Model(&models.User{}).Preload("UserDetail.Adress.District.Province").Where("phone = ?", user.Phone).Take(&user).Error
	return err
}

// register new user..
func (u UserRepo) Register(user models.User) error {
	err := u.db.Model(&models.User{}).Save(&user)
	return err.Error
}

func (u UserRepo) Update(user models.User) error {
	err := u.db.Save(&user).Error
	return err
}

// exists phone ?
func (u UserRepo) ExistsPhone(phone string, userid uint) bool {
	err := u.db.Model(&models.User{}).Where("phone = ?", phone).Where("id != ?", userid).Find(&models.User{})
	if err.RowsAffected > 0 {
		return true
	}
	return false
}

func (u UserRepo) ExistsEmail(email string, userid uint) bool {
	err := u.db.Model(&models.User{}).Where("email = ?", email).Where("id != ?", userid).Find(&models.User{})
	if err.Error != nil {
		return true
	} else if err.RowsAffected > 0 {
		return true
	}
	return false
}
