package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string         `json:"id" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
	Name       string         `json:"name" form:"name"`
	Email      string         `json:"email" form:"email" gorm:"not null;unique;size:20;"`
	Password   string         `json:"password" form:"password" gorm:"not null"`
	Age        int            `json:"age" form:"age" gorm:"default:18"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

