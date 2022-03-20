package model

import "gorm.io/gorm"

type Promo struct {
	gorm.Model
	CompanyName string   `json:"company_name"`
	Phone       int      `json:"phone"`
	Photo       Photo    `gorm:"foreignkey:photofk" json:"photofk"`
	Photofk     uint     `gorm:"column:photofk"`
	Category    Category `gorm:"foreignkey:categoryfk" json:"category"`
	Categoryfk  uint     `gorm:"column:categoryfk" json:"categoryfk"`
}

type PromoRequest struct {
	gorm.Model
	CompanyName string   `json:"company_name"`
	Phone       int      `json:"phone"`
	Photo       Photo    `gorm:"foreignkey:photofk" json:"photofk"`
	Photofk     uint     `gorm:"column:photofk"`
	Category    Category `gorm:"foreignkey:categoryfk" json:"category"`
	Categoryfk  uint     `gorm:"column:categoryfk" json:"categoryfk"`
}
