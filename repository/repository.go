package repository

import (
	config "ilanver/internal/configs"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	Db *gorm.DB
}

var Repo *Repositories

func Set() {
	Repo = &Repositories{
		Db: config.Database.
			Set("gorm:association_autoupdate", false).
			Set("gorm:association_autocreate", false).
			Set("gorm:association_save_reference", false),
	}
}

func Get() *Repositories {
	return Repo
}
