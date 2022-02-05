package repository

import (
	"ilanver/internal/models"

	"github.com/jinzhu/gorm"
)

func (repo *Repositories) Address() AddressRepo {
	return AddressRepo{db: repo.Db}
}

type AddressRepo struct {
	db *gorm.DB
}

func (a AddressRepo) Save(adres *models.Adress) error {
	err := a.db.Model(&models.Adress{}).Save(&adres)
	return err.Error
}

func (a AddressRepo) ExistsDistric(district uint) bool {
	err := a.db.Model(&models.District{}).Where("id = ?", district)
	if err.Error != nil {
		return true
	} else if err.RowsAffected > 0 {
		return true
	}
	return false
}
