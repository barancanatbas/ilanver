package model

import (
	"database/sql"
	"time"
)

type Adress struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	Detail     string       `json:"detail"`
	Districtfk uint         `gorm:"column:districtfk" json:"districtfk"`
	District   District     `gorm:"foreignkey:districtfk"`
}

type District struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	District   string       `json:"district"`
	Provincefk uint         `gorm:"column:provincefk" json:"province"`
	Province   Province     `gorm:"foreignkey:provincefk"`
}

type Province struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Province  string
}
