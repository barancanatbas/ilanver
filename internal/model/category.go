package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string
	MainCategory uint `gorm:"column:maincategory;default:null"`
	Src          string
}
