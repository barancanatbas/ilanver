package repository

import (
	"ilanver/internal/config"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateTX() *gorm.DB
	SetTX(tx *gorm.DB)
	RollBack()
	Commit()
}

type Repository struct {
	tx *gorm.DB
}

func NewRepository(tx *gorm.DB) IRepository {
	return &Repository{
		tx: tx,
	}
}

func (r *Repository) CreateTX() *gorm.DB {
	r.tx = r.tx.Begin()
	return r.tx
}

func (r *Repository) SetTX(tx *gorm.DB) {
	r.tx = tx
}

func (r *Repository) RollBack() {
	r.tx.Rollback()
}

func (r *Repository) Commit() {
	r.tx.Commit()

	r.tx = config.DB
}
