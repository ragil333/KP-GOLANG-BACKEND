package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	BookId    uint           `json:"book_id" gorm:"primaryKey;type:int"`
	Title     string         `json:"title" validate:"required"`
	Author    string         `json:"author" validate:"required"`
	Publisher string         `json:"publisher" validate:"required"`
	Year      int            `json:"year" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
