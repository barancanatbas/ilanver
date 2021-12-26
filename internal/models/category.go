package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryName string
	MainCategory uint `gorm:"column:maincategory;default:null"`
	Src          string
}
