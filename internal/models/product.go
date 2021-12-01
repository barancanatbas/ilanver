package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title  string
	User   User `gorm:"column:userfk" json:"user"`
	Userfk uint `gorm:"foreignkey:userfk" json:"userfk"`
}

type ProductDetail struct {
	gorm.Model
	Adress         Adress       `gorm:"column:adressfk" json:"adress"`
	Adressfk       uint         `gorm:"foreignkey:adressfk" json:"adressfk"`
	ProductState   ProductState `gorm:"column:product_statefk" json:"prodoct_state"`
	ProductStatefk uint         `gorm:"foreignkey:product_statefk" json:"product_statefk"` // satılık kiralık vs diye
	Description    string       `json:"description"`
	Price          uint         `json:"price"`
	Category       Category     `gorm:"column:categoryfk" json:"category"`
	Categoryfk     uint         `gorm:"foreignkey:categoryfk" json:"categoryfk"`
}

type ProductState struct {
	gorm.Model
	State string
}
