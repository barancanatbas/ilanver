package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string     `json:"name" gorm:"type:varchar(40)"`
	Surname      string     `json:"surname" gorm:"type:varchar(40)"`
	Phone        string     `json:"phone" gorm:"unique;not null;type:char(11)"`
	Password     string     `json:"-"`
	Email        string     `json:"email"`
	UserDetail   UserDetail `gorm:"foreignkey:userdetailfk" json:"user_detail"`
	UserDetailfk uint       `gorm:"column:userdetailfk" json:"user_detailfk"`
	Birthday     time.Time  `json:"birthday"`
}

type UserDetail struct {
	gorm.Model
	ProfilePhoto string `json:"profile_photo"`
	Adressfk     uint   `gorm:"column:adressfk"`
	Adress       Adress `gorm:"foreignkey:adressfk" json:"adressfk"`
}

type LostPassword struct {
	gorm.Model
	User   User `gorm:"foreignkey:userfk" json:"user"`
	Userfk uint `gorm:"column:userfk" json:"userfk"`
	Code   string
}
