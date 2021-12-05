package models

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Src       string
	Product   Product `gorm:"foreignkey:productfk" json:"product"`
	Productfk uint    `gorm:"column:productfk" json:"productfk"`
}
