package models

import "github.com/jinzhu/gorm"

type LostPassword struct {
	gorm.Model
	User   User `gorm:"foreignkey:userfk" json:"user"`
	Userfk uint `column:userfk" json:"userfk"`
}
