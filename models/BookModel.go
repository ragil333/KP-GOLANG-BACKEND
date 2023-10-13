package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	BookId    uint           `json:"book_id" gorm:"primaryKey;type:int"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Publisher string         `json:"publisher"`
	Year      int            `json:"year"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
