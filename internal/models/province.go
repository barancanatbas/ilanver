package models

import "github.com/jinzhu/gorm"

type Province struct {
	gorm.Model
	Province string
}
