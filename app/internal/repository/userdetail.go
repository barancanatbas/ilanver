package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IUserDetailRepo interface {
	Save(detail *model.UserDetail) error
	GetByID(id uint) (model.UserDetail, error)
	WithTx(db *gorm.DB) IUserDetailRepo
}

type UserDetailRepo struct {
	tx *gorm.DB
}

func NewUserDetailRepository(tx *gorm.DB) IUserDetailRepo {
	return &UserDetailRepo{
		tx: tx,
	}
}

func (u UserDetailRepo) WithTx(db *gorm.DB) IUserDetailRepo {
	u.tx = db

	return u
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
