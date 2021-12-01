package models

import "github.com/jinzhu/gorm"

type Promo struct {
	gorm.Model
	CompanyName string   `json:"company_name"`
	Phone       int      `json:"phone"`
	Photo       Photo    `gorm:"column:photofk" json:"photofk"`
	Photofk     uint     `gorm:"foreignkey:photofk"`
	Category    Category `gorm:"column:categoryfk" json:"category"`
	Categoryfk  uint     `gorm:"foreignkey:categoryfk" json:"categoryfk"`
}

type PromoRequest struct {
	gorm.Model
	CompanyName string   `json:"company_name"`
	Phone       int      `json:"phone"`
	Photo       Photo    `gorm:"column:photofk" json:"photofk"`
	Photofk     uint     `gorm:"foreignkey:photofk"`
	Category    Category `gorm:"column:categoryfk" json:"category"`
	Categoryfk  uint     `gorm:"foreignkey:categoryfk" json:"categoryfk"`
}
