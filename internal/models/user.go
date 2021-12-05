package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Phone        string     `json:"phone"`
	Password     string     `json:"-"`
	Email        string     `json:"email"`
	UserDetail   UserDetail `gorm:"column:userdetailfk" json:"user_detail"`
	UserDetailfk uint       `gorm:"foreignkey:userdetailfk" json:"user_detailfk"`
}

type UserDetail struct {
	gorm.Model
	ProfilePhoto string `json:"profile_photo"`
	Adressfk     uint   `gorm:"foreignkey:adressfk" json:"adressfk"`
	Adress       Adress `gorm:"column:adressfk"`
	Birthday     string `json:"birthday"`
}
