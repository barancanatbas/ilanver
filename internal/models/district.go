package models

import "github.com/jinzhu/gorm"

type District struct {
	gorm.Model
	District   string   `json:"district"`
	Provincefk uint     `gorm:"column:provincefk" json:"province"`
	Province   Province `gorm:"foreignkey:provincefk"`
}
