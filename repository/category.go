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

func (ct CategoryRepo) Exists(id uint) (models.Category, error) {
	category := models.Category{}
	val := ct.db.Model(&models.Category{}).Where("id = ?", id).Take(&category)
	if val.RowsAffected > 0 {
		return category, val.Error
	}
	return category, nil
}

func (ct CategoryRepo) ExistsMain(main_category_id uint) (models.Category, error) {
	category := models.Category{}
	err := ct.db.Model(&models.Category{}).
		Where("main_category = ?", main_category_id).Find(&category).Error
	return category, err
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

func (ct CategoryRepo) Update(category *models.Category) error {
	err := ct.db.Save(&category).Error
	return err
}
