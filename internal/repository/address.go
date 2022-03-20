package repository

import (
	"ilanver/internal/model"

	"gorm.io/gorm"
)

type IAddressRepo interface {
	Save(address *model.Adress) error
}

type AddressRepo struct {
	tx *gorm.DB
}

// Compile time checks to ensure your type satisfies an interface
var _ IAddressRepo = AddressRepo{}

func NewAddressRepository(tx *gorm.DB) AddressRepo {
	return AddressRepo{
		tx: tx,
	}
}

func (a AddressRepo) Save(address *model.Adress) error {
	err := a.tx.Save(address).Error

	return err
}
