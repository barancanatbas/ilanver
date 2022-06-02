package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetSubCategories(id uint) ([]model.Category, error)
	Insert(category *model.Category) error
	GetByID(id int) (model.Category, error)
	Update(category model.Category) error
	Delete(id uint) error
	DeleteWithInQuery(data []int) error
	//Paginate(page int, status int) (helpers.Paginate, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

// func (c CategoryRepository) GetAll(page int) (helpers.Pagination, error) {
// 	// only get main category
// 	var categories []model.Category
// 	pagination := helpers.Pagination{
// 		Limit: 10,
// 		Page:  page,
// 	}

// 	c.db = c.db.Debug().Scopes(helpers.Paginate(categories, &pagination, c.db)).
// 		Table("categories").
// 		Where("maincategory = ? or maincategory is null", 0).
// 		Find(&categories)
// 	pagination.Rows = categories

// 	return pagination, nil
// }

func (c CategoryRepository) GetAll() ([]model.Category, error) {
	// only get main category
	var categories []model.Category

	c.db = c.db.Table("categories").Where("maincategory = ? or maincategory is null", 0).Find(&categories)

	return categories, nil
}

func (c CategoryRepository) GetSubCategories(id uint) ([]model.Category, error) {
	var categories []model.Category
	err := c.db.Where("maincategory = ?", id).Find(&categories).Error
	return categories, err
}

func (c CategoryRepository) GetByID(id int) (model.Category, error) {
	var category model.Category
	err := c.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (c CategoryRepository) Insert(category *model.Category) error {
	return c.db.Create(category).Error
}

func (c CategoryRepository) Update(category model.Category) error {
	return c.db.Save(&category).Error
}

func (c CategoryRepository) Delete(id uint) error {
	return c.db.Delete(&model.Category{}, id).Error
}

func (c CategoryRepository) DeleteWithInQuery(data []int) error {
	return c.db.Where("id IN (?)", data).Delete(&model.Category{}).Error
}
