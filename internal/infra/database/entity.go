package database

import (
	"time"

	"gorm.io/gorm"
)

type (
	Entity interface {
		GetID() uint
	}

	BaseEntity struct {
		ID        uint           `json:"id" param:"id" query:"id" gorm:"primary_key"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	}
)

var _ Entity = (*BaseEntity)(nil)

func (e BaseEntity) GetID() uint {
	return e.ID
}
