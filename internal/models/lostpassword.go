package models

import "github.com/jinzhu/gorm"

type LostPassword struct {
	gorm.Model
	User   User `gorm:"column:userfk" json:"user"`
	Userfk uint `gorm:"foreignkey:userfk" json:"userfk"`
}
