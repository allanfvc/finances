package model

import (
	"github.com/jinzhu/gorm"
)

type Input struct {
	gorm.Model
	Description string  `json:"description" gorm:"not null"`
	Value       float32 `json:"value" gorm:"not null"`
	TypeInput   int     `json:"type" gorm:"not null"`
	Status      int     `json:"status" gorm:"not null"`
	User        User    `json:"user" gorm:"not null"`
	Month       int     `json:"month" gorm:"not null"`
	Year        int     `json:"year" gorm:"not null"`
}
