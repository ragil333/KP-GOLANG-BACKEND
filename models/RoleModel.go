package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	RoleId    uint           `json:"role_id" gorm:"primaryKey;type:int"`
	Role      string         `json:"role" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
