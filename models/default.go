package models

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	ID uint `gorm:"primarykey" json:"id"`
}

type Timestamp struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
