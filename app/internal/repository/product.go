package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetByID(id int) (model.Product, error)
	Save(product *model.Product) error
	SaveDetail(product *model.ProductDetail) error
	GetDetail(id uint, productD *model.ProductDetail) error
	WithTx(db *gorm.DB) IProductRepository
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (a *ProductRepository) WithTx(db *gorm.DB) IProductRepository {
	a.db = db

	return a
}

func (p *ProductRepository) GetByID(id int) (model.Product, error) {
	var product model.Product
	err := p.db.Preload("ProductDetail").Where("id = ?", id).First(&product).Error
	return product, err
}

func (p *ProductRepository) SaveDetail(product *model.ProductDetail) error {
	return p.db.Save(product).Error
}

func (p *ProductRepository) Save(product *model.Product) error {
	err := p.db.Save(product).Error
	return err
}

func (p *ProductRepository) GetDetail(id uint, productD *model.ProductDetail) error {
	err := p.db.Where("id = ?", id).First(productD).Error
	return err
}
