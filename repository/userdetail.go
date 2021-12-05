package repository

import (
	"ilanver/internal/models"

	"github.com/jinzhu/gorm"
)

type userdetail struct {
	db *gorm.DB
}

func (repo *Repositories) UserDetail() userdetail {
	return userdetail{db: repo.Db}
}

func (ud userdetail) Save(detail *models.UserDetail) error {
	err := ud.db.Save(&detail)
	return err.Error
}
