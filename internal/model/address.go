package model

import "gorm.io/gorm"

type Adress struct {
	gorm.Model
	Detail     string   `json:"detail"`
	Districtfk uint     `gorm:"column:districtfk" json:"districtfk"`
	District   District `gorm:"foreignkey:districtfk"`
}

type District struct {
	gorm.Model
	District   string   `json:"district"`
	Provincefk uint     `gorm:"column:provincefk" json:"province"`
	Province   Province `gorm:"foreignkey:provincefk"`
}

type Province struct {
	gorm.Model
	Province string
}
