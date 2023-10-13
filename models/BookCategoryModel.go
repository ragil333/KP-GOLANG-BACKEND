package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookCategory struct {
	BookCategoryId uuid.UUID `json:"book_category_id" gorm:"primaryKey;type:uuid"`
	CategoryId     uuid.UUID `json:"category_id"`
	// Category       Category       `json:"category" gorm:"foreignKey:CategoryId;contraint:OnUpdate:Cascade,OnDelete:Cascade"`
	BookId uuid.UUID `json:"book_id"`
	// Book           Book           `json:"book" gorm:"foreignKey:BookId;contraint:OnUpdate:Cascade,OnDelete:Cascade"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
