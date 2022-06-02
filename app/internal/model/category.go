package model

import (
	"database/sql"
	"time"
)

type Category struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
	CategoryName string
	MainCategory uint `gorm:"column:maincategory;default:null"`
	Src          string
}
