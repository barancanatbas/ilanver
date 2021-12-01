package models

import "github.com/jinzhu/gorm"

type Adress struct {
	gorm.Model
	Detail     string   `json:"detail"`
	Districtfk uint     `gorm:"foreignkey:districtfk" json:"districtfk"`
	District   District `gorm:"column:districtfk"`
}
