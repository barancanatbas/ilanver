package models

import "github.com/jinzhu/gorm"

type District struct {
	gorm.Model
	District   string   `json:"district"`
	Provincefk uint     `gorm:"foreignkey:provincefk"`
	Province   Province `gorm:"column:provincefk" json:"province"`
}
