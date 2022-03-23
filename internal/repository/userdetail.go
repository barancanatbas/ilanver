package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IUserDetailRepo interface {
	Save(detail *model.UserDetail) error
	GetByID(id uint) (model.UserDetail, error)
}

type UserDetailRepo struct {
	tx *gorm.DB
}

// Compile time checks to ensure your type satisfies an interface
var _ IUserDetailRepo = UserDetailRepo{}

func NewUserDetailRepository(tx *gorm.DB) UserDetailRepo {
	return UserDetailRepo{
		tx: tx,
	}
}

func (u UserDetailRepo) Save(detail *model.UserDetail) error {

	err := u.tx.Save(detail).Error

	return err
}

func (u UserDetailRepo) GetByID(id uint) (model.UserDetail, error) {
	var detail model.UserDetail

	err := u.tx.Model(&model.UserDetail{}).Where("id = ?", id).First(&detail).Error

	return detail, err
}
