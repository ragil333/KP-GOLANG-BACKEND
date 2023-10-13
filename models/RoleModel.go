package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	RoleId    uint           `json:"role_id" gorm:"primaryKey;type:int"`
	Role      string         `json:"role" validate:"required"`
	User      []User         `json:"user" gorm:"Foreignkey:RoleId;association_foreignkey:UserId;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
