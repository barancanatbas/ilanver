package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IProductStateRepository interface {
	GetByID(id int) (model.ProductState, error)
	GetAll() ([]model.ProductState, error)
	Insert(productState *model.ProductState) error
	Update(productState *model.ProductState) error
	Delete(id uint) error
}

type ProductStateRepository struct {
	db *gorm.DB
}

func NewProductStateRepository(db *gorm.DB) IProductStateRepository {
	return &ProductStateRepository{
		db: db,
	}
}

func (p *ProductStateRepository) GetByID(id int) (model.ProductState, error) {
	var productState model.ProductState
	err := p.db.Where("id = ?", id).First(&productState).Error
	return productState, err
}

func (p *ProductStateRepository) GetAll() ([]model.ProductState, error) {
	var productStates []model.ProductState
	err := p.db.Find(&productStates).Error
	return productStates, err
}

func (p *ProductStateRepository) Insert(productState *model.ProductState) error {
	err := p.db.Create(productState).Error
	return err
}

func (p *ProductStateRepository) Update(productState *model.ProductState) error {
	err := p.db.Save(productState).Error
	return err
}

func (p *ProductStateRepository) Delete(id uint) error {
	err := p.db.Where("id = ?", id).Delete(&model.ProductState{}).Error
	return err
}
