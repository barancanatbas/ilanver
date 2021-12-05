package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryName string
	MainCategory uint
	Src          string
}
