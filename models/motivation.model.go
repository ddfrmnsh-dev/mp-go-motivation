package models

import (
	"time"

	"gorm.io/gorm"
)

type Motivation struct {
	Id        int            `json:"id" gorm:"primaryKey"`
	Quote     string         `json:"quote" form:"quote" validate:"gte=6,lte=32" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
