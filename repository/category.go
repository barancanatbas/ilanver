package repository

import (
	"ilanver/internal/models"

	"github.com/jinzhu/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func (repo *Repositories) Category() CategoryRepo {
	return CategoryRepo{db: repo.Db}
}

func (ct CategoryRepo) Exists(id uint) bool {
	val := ct.db.Model(&models.Category{}).Where("id = ?", id).Take(&models.Category{})
	if val.RowsAffected > 0 {
		return true
	}
	return false
}

func (ct CategoryRepo) Insert(category *models.Category) error {
	val := ct.db.Model(&models.Category{}).Save(&category)

	return val.Error
}

func (ct CategoryRepo) MainCategory() ([]models.Category, error) {
	categorys := []models.Category{}
	val := ct.db.Model(&models.Category{}).Where("main_category = ?", 0).Find(&categorys)

	return categorys, val.Error
}
