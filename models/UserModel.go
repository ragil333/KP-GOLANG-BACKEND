package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	UserId   uuid.UUID `json:"user_id" gorm:"primarykey;type:uuid"`
	Name     string    `json:"name" gorm:"not null" validate:"required"`
	Email    string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string    `json:"password" gorm:"not null" validate:"required,min=8"`
	RoleId   uuid.UUID `json:"role_id"`
	// Role      Role           `json:"role" gorm:"foreignKey:RoleId;contraint:OnUpdate:Cascade,OnDelete:SET NULL"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
