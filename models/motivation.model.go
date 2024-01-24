package models

import (
	"time"
)

type Motivation struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Quote     string    `json:"quote" form:"quote" gorm:"not null"`
	Name      string    `json:"name" form:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
