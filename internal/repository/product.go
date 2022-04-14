package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetByID(id int) (model.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p ProductRepository) GetByID(id int) (model.Product, error) {
	var product model.Product
	err := p.db.Preload("ProductDetail").Where("id = ?", id).First(&product).Error
	return product, err
}
