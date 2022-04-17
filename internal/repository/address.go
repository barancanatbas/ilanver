package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IAddressRepo interface {
	Save(address *model.Adress) error
	GetByID(id uint) (model.Adress, error)
	WithTx(db *gorm.DB) IAddressRepo
	Update(address *model.Adress) error
}

type AddressRepo struct {
	tx *gorm.DB
}

func NewAddressRepository(tx *gorm.DB) IAddressRepo {
	return &AddressRepo{
		tx: tx,
	}
}

func (a AddressRepo) WithTx(db *gorm.DB) IAddressRepo {
	a.tx = db

	return a
}

func (a AddressRepo) Save(address *model.Adress) error {
	err := a.tx.Save(address).Error

	return err
}

func (a AddressRepo) GetByID(id uint) (model.Adress, error) {
	var address model.Adress

	err := a.tx.Model(&model.Adress{}).Where("id = ?", id).First(&address).Error

	return address, err
}

func (a AddressRepo) Update(address *model.Adress) error {
	err := a.tx.Save(address).Error

	return err
}
