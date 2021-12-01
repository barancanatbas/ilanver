package models

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Src       string
	Product   Product `gorm:"column:productfk" json:"product"`
	Productfk uint    `gorm:"foreignkey:productfk" json:"productfk"`
}
