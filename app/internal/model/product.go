package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title           string
	User            User          `gorm:"foreignkey:userfk" json:"user"`
	Userfk          uint          `gorm:"column:userfk" json:"userfk"`
	ProductDetail   ProductDetail `gorm:"foreignkey:productdetailfk" json:"productdetail"`
	ProductDetailfk uint          `gorm:"column:productdetailfk" json:"productdetailfk"`
}

type ProductDetail struct {
	gorm.Model
	Adress         Adress       `gorm:"foreignkey:adressfk" json:"adress"`
	Adressfk       uint         `gorm:"column:adressfk" json:"adressfk"`
	ProductState   ProductState `gorm:"foreignkey:product_statefk" json:"prodoct_state"`
	ProductStatefk uint         `gorm:"column:product_statefk" json:"product_statefk"` // satılık kiralık vs diye
	Description    string       `json:"description"`
	Price          uint         `json:"price"`
	Category       Category     `gorm:"foreignkey:categoryfk" json:"category"`
	Categoryfk     uint         `gorm:"column:categoryfk" json:"categoryfk"`
}

type ProductState struct {
	gorm.Model
	State string
}

type ProductElastic struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Category    string `json:"category"`
}
