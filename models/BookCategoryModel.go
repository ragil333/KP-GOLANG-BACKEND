package models

import (
	"time"

	"gorm.io/gorm"
)

type BookCategory struct {
	BookCategoryId uint `json:"book_category_id" gorm:"primaryKey;type:int"`
	CategoryId     uint `json:"category_id"`
	// Category       Category       `json:"category" gorm:"foreignKey:CategoryId;contraint:OnUpdate:Cascade,OnDelete:Cascade"`
	BookId uint `json:"book_id"`
	// Book           Book           `json:"book" gorm:"foreignKey:BookId;contraint:OnUpdate:Cascade,OnDelete:Cascade"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
