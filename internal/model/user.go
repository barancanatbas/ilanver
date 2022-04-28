package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
	Name         string       `json:"name" gorm:"type:varchar(40)"`
	Surname      string       `json:"surname" gorm:"type:varchar(40)"`
	Phone        string       `json:"phone" gorm:"unique;not null;type:char(11)"`
	Password     string       `json:"-"`
	Email        string       `json:"email"`
	UserDetail   UserDetail   `gorm:"foreignkey:userdetailfk" json:"user_detail"`
	UserDetailfk uint         `gorm:"column:userdetailfk" json:"user_detailfk"`
	Birthday     time.Time    `json:"birthday"`
}

type UserDetail struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
	ProfilePhoto string       `json:"profile_photo"`
	Adressfk     uint         `gorm:"column:adressfk"`
	Adress       Adress       `gorm:"foreignkey:adressfk" json:"adressfk"`
}

type LostPassword struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	User      User         `gorm:"foreignkey:userfk" json:"user"`
	Userfk    uint         `gorm:"column:userfk" json:"userfk"`
	Code      string
}
