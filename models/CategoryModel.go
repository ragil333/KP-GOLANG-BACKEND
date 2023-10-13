package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CategoryId uint           `json:"category_id" gorm:"primaryKey;type:int"`
	Category   string         `json:"category" validate:"required"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
